package handler

import (
	. "middleware"
	//"net/http"
	"log"
	"model"
	"regexp"
	"strings"
)


type ChatHandler struct {
	Huber
	Middler
}

//处理客户端对websocket请求
// serveWs handles websocket requests from the peer.

func (x *ChatHandler)Get() {

	//设定环境变量
	ws, err := Upgrader.Upgrade(x.Ctx.ResponseWriter, x.Ctx.Req(), nil)
	if err != nil {
		log.Println("error:",err)
		return
	}
	//初始化Connection
	c := &Connection{Send: make(chan []byte, 256), Ws: ws, Auth: false,}

	//验证用户
	chatuser_session := x.Session.Get(SESSION_VALUE_USERLOGIN)

	if chatuser_session!=nil{
		chatuser := chatuser_session.(model.ChatUser)

		log.Println(chatuser,chatuser.Username)

		groupid 	:= []byte("1")
		username 	:= []byte(chatuser.Username)
		fromuser 	:= []byte(chatuser.Username)
		touser 		:= []byte("all")
		mtype 		:= 3

		c = &Connection{Send: make(chan []byte, 256), Ws: ws, Auth: false, Mtype:mtype, Username:username, Groupid:groupid, Fromuser:fromuser, Touser:touser,}

		//log.Println(c)
		//加入注册通道，意思是只要连接的人都加入register通道
		Hub.Register <- c
		go c.WritePump() //服务器端发送消息给客户端
		c.ReadPump()     //服务器读取的所有客户端的发来的消息
	}else{

		c.Send <- []byte("nologin")
		go c.WritePump() //服务器端发送消息给客户端
		c.ReadPump()     //服务器读取的所有客户端的发来的消息
	}
}



func ChatHandlerRun() {

	for {
		select {
		//注册者有数据，则插入连接map
		case c := <-Hub.Register:
			Hub.Connections[c] = true

			//发送通知数据给连接
			for d := range Hub.Connections {
				var send_msg []byte
				send_msg = []byte(string("欢迎" + string(c.Username) + "加入聊天"))

				select {
				//发送数据给连接
				case d.Send <- send_msg:
				//关闭连接
				default:
					close(d.Send)
					delete(Hub.Connections, d)
				}
			}

		//非注册者有数据，则删除连接map
		case c := <-Hub.Unregister:

			//发送通知数据给连接
			for d := range Hub.Connections {
				var send_msg []byte
				send_msg = []byte(string( string(c.Username) + "离开了聊天室"))

				select {
				//发送数据给连接
				case d.Send <- send_msg:
				//关闭连接
				default:
					close(d.Send)
					delete(Hub.Connections, d)
				}
			}

			if _, ok := Hub.Connections[c]; ok {
				delete(Hub.Connections, c)
				close(c.Send)
			}

		//广播有数据
		case m := <-Hub.Broadcast:
		//递归所有广播连接
			for c := range Hub.Connections {
				var send_flag = false

				//根据广播消息标识记录
				/*
					text2 := string(m.content)
					reg2 := regexp.MustCompile(`^@.*? `)
					s2 := reg2.FindAllString(text2, -1)
				*/
				//查找@标签　
				check_content := string(m.Content)
				find_flag := regexp.MustCompile(`^@.*? `)
				find_user := find_flag.FindAllString(check_content,-1)

				log.Println(find_user)

				if find_user != nil {

					m.Mtype = 2
					find_user[0] = strings.Replace(find_user[0], "@", "", 1)
					find_user[0] = strings.Replace(find_user[0], " ", "", 1)
					m.Touser = []byte(find_user[0])
				}

				var send_msg []byte
				if m.Mtype == 1 { //系统消息
					send_msg = []byte(" system: " + string(m.Content))
				} else if m.Mtype == 2 { //用户消息
					send_msg = []byte(string(m.Fromuser) + " say to you: " + string(m.Content))
				} else {
					send_msg = []byte(string(m.Fromuser) + " say: " + string(m.Content))
				}

				log.Print(send_msg,"sess::")
				if string(m.Touser) != "all" {
					if string(c.Username) == string(m.Touser) || string(c.Username) == string(m.Fromuser) {
						send_flag = true
					}
					if send_flag {
						select {
						//发送数据给连接
						case c.Send <- send_msg:
						//关闭连接
						default:
							close(c.Send)
							delete(Hub.Connections, c)
						}
					}
				} else {
					select {
					//发送数据给连接
					case c.Send <- send_msg:
					//关闭连接
					default:
						close(c.Send)
						delete(Hub.Connections, c)
					}
				}

			}
		}
	}
}

