package middleware


import (
	"github.com/gorilla/websocket"
	"time"
	"log"
)

const (
	//对方写入会话等待时间
	// Time allowed to write a message to the peer.
	WriteWait = 10 * time.Second

	//对方读取下次消息等待时间
	// Time allowed to read the next pong message from the peer.
	PongWait = 60 * time.Second

	//对方ping周期
	// Send pings to peer with this period. Must be less than pongWait.
	PingPeriod = (PongWait * 9) / 10

	//对方最大写入字节数
	// Maximum message size allowed from peer.
	MaxMessageSize = 512

	//验证字符串
	AuthToken = "123456"
)

//服务器配置信息
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Connection 是websocket的conntion和hub的中间人
// Connection is an middleman between the websocket connection and the hub.
type Connection struct {
	// The websocket Connection.
	//websocket的连接
	Ws *websocket.Conn

	// Buffered channel of outbound messages.
	//出站消息缓存通道
	Send chan []byte

	//验证状态
	Auth bool

	//消息类型
	Mtype int

	//验证状态
	Username []byte

	//分组标示
	Groupid []byte

	//来自哪个用户
	Fromuser []byte

	//对哪个用户
	Touser []byte


}



//读取Connection中的数据导入到hub中，实则发广播消息
//服务器读取的所有客户端的发来的消息
// readPump pumps messages from the websocket Connection to the hub.
func (c *Connection) ReadPump()  {
	defer func() {
		//h.unregister <- c
		c.Ws.Close()
	}()
	c.Ws.SetReadLimit(MaxMessageSize)
	c.Ws.SetReadDeadline(time.Now().Add(PongWait))
	c.Ws.SetPongHandler(func(string) error { c.Ws.SetReadDeadline(time.Now().Add(PongWait)); return nil })

	for {
		_, message, err := c.Ws.ReadMessage()
		if err != nil {
			break
		}


		t := time.Now().Unix()
		Hub.Broadcast <- &Tmessage{Content: message, Createtime: time.Unix(t, 0).String(), Mtype:c.Mtype, Fromuser:c.Fromuser, Touser:c.Touser, Groupid:c.Groupid}
		log.Println(string(message),  string(c.Username),   time.Unix(t, 0).String(),string(c.Fromuser),string(c.Touser),string(c.Groupid))

	}

}


//给消息，指定消息类型和荷载
// write writes a message with the given message type and payload.
func (c *Connection) Write(mt int, payload []byte) error {
	c.Ws.SetWriteDeadline(time.Now().Add(WriteWait))
	return c.Ws.WriteMessage(mt, payload)
}

//从hub到Connection写数据
//服务器端发送消息给客户端
// writePump pumps messages from the hub to the websocket Connection.
func (c *Connection) WritePump() {
	//定时执行
	ticker := time.NewTicker(PingPeriod)

	defer func() {
		ticker.Stop()
		c.Ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:

			if !ok {
				log.Println("conn_writeump_case_1")
				c.Write(websocket.CloseMessage, []byte{})

				return
			}

			if err := c.Write(websocket.TextMessage, message); err != nil {
				log.Println("conn_writeump_case_2")
				return
			}
		case <-ticker.C:
			if err := c.Write(websocket.PingMessage, []byte{}); err != nil {
				log.Println("conn_writeump_case_3")
				return
			}
		}
	}

	//test

}




type Tmessage struct {
	Content    []byte
	Fromuser   []byte
	Touser     []byte
	Mtype      int
	Createtime string
	Groupid	   []byte
}

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type Huber struct {
	// Registered connections.
	//注册连接
	Connections map[*Connection]bool

	// Inbound messages from the connections.
	//连接中的绑定消息
	Broadcast chan *Tmessage

	// Register requests from the connections.
	//添加新连接
	Register chan *Connection

	// Unregister requests from connections.
	//删除连接
	Unregister chan *Connection
}

var Hub = Huber{
	//广播slice
	Broadcast: make(chan *Tmessage),
	//注册者slice
	Register: make(chan *Connection),
	//未注册者sclie
	Unregister: make(chan *Connection),
	//连接map
	Connections: make(map[*Connection]bool),
}
