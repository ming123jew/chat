package middleware


import (
	"github.com/gorilla/websocket"
	"time"
	"regexp"
	"strings"
	"log"
	"fmt"
	"net/http"
)

const (
	//对方写入会话等待时间
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	//对方读取下次消息等待时间
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	//对方ping周期
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	//对方最大写入字节数
	// Maximum message size allowed from peer.
	maxMessageSize = 512

	//验证字符串
	authToken = "123456"
)

//服务器配置信息
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Connection 是websocket的conntion和hub的中间人
// Connection is an middleman between the websocket connection and the hub.
type Connection struct {
	// The websocket Connection.
	//websocket的连接
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	//出站消息缓存通道
	send chan []byte

	//验证状态
	auth bool

	//验证状态
	username []byte

	C
}



//读取Connection中的数据导入到hub中，实则发广播消息
//服务器读取的所有客户端的发来的消息
// readPump pumps messages from the websocket Connection to the hub.
func (c *Connection) readPump() {
	defer func() {
		//h.unregister <- c
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}

		mtype := 2 //用户消息
		text := string(message)
		log.Println(text)
		reg := regexp.MustCompile(`=[^&]+`)
		log.Println(reg)
		s := reg.FindAllString(text, -1)
		log.Println(s)

		//默认all
		if len(s) == 2 {
			fromuser := strings.Replace(s[0], "=", "", 1)
			token := strings.Replace(s[1], "=", "", 1)
			if token == authToken {
				c.username = []byte(fromuser)
				c.auth = true
				message = []byte(fromuser + " join")
				mtype = 1 //系统消息
			}
		}

		touser := []byte("all")
		reg2 := regexp.MustCompile(`^@.*? `)
		s2 := reg2.FindAllString(text, -1)
		if len(s2) == 1 {
			s2[0] = strings.Replace(s2[0], "@", "", 1)
			s2[0] = strings.Replace(s2[0], " ", "", 1)
			touser = []byte(s2[0])
		}
		t := time.Now().Unix()
		HUB.broadcast <- &Tmessage{content: message, fromuser: c.username, touser: touser, mtype: mtype, createtime: time.Unix(t, 0).String()}
		log.Println(message,  c.username, touser, mtype,  time.Unix(t, 0).String())
		if c.auth == true {

		}
	}
}


//给消息，指定消息类型和荷载
// write writes a message with the given message type and payload.
func (c *Connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

//从hub到Connection写数据
//服务器端发送消息给客户端
// writePump pumps messages from the hub to the websocket Connection.
func (c *Connection) writePump() {
	//定时执行
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:

			if !ok {
				fmt.Println("conn_writeump_case_1")
				c.write(websocket.CloseMessage, []byte{})

				return
			}

			if err := c.write(websocket.TextMessage, message); err != nil {
				fmt.Println("conn_writeump_case_2")
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				fmt.Println("conn_writeump_case_3")
				return
			}
		}
	}

	//test

}


//处理客户端对websocket请求
// serveWs handles websocket requests from the peer.

func ServeWs(w http.ResponseWriter, r *http.Request) {

	//设定环境变量
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}


	//初始化Connection
	c := &Connection{send: make(chan []byte, 256), ws: ws, auth: false}
	log.Println(c)
	//加入注册通道，意思是只要连接的人都加入register通道
	HUB.register <- c
	go c.writePump() //服务器端发送消息给客户端
	c.readPump()     //服务器读取的所有客户端的发来的消息
}