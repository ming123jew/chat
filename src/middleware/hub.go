package middleware

// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
//"fmt"
)

type Tmessage struct {
	content    []byte
	fromuser   []byte
	touser     []byte
	mtype      int
	createtime string
}

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	// Registered connections.
	//注册连接
	connections map[*Connection]bool

	// Inbound messages from the connections.
	//连接中的绑定消息
	broadcast chan *Tmessage

	// Register requests from the connections.
	//添加新连接
	register chan *Connection

	// Unregister requests from connections.
	//删除连接
	unregister chan *Connection
}

var HUB = Hub{
	//广播slice
	broadcast: make(chan *Tmessage),
	//注册者slice
	register: make(chan *Connection),
	//未注册者sclie
	unregister: make(chan *Connection),
	//连接map
	connections: make(map[*Connection]bool),
}

func (h *Hub) Run() {
	for {
		select {
		//注册者有数据，则插入连接map
		case c := <-h.register:
			h.connections[c] = true

		//非注册者有数据，则删除连接map
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}

		//广播有数据
		case m := <-h.broadcast:
		//递归所有广播连接
			for c := range h.connections {
				var send_flag = false

				//根据广播消息标识记录
				/*
					text2 := string(m.content)
					reg2 := regexp.MustCompile(`^@.*? `)
					s2 := reg2.FindAllString(text2, -1)
				*/
				var send_msg []byte
				if m.mtype == 1 { //系统消息
					send_msg = []byte(" system: " + string(m.content))
				} else if m.mtype == 2 { //用户消息
					send_msg = []byte(string(m.fromuser) + " say: " + string(m.content))
				} else {
					send_msg = []byte(string(m.content))
				}
				if string(m.touser) != "all" {
					if string(c.username) == string(m.touser) || string(c.username) == string(m.fromuser) {
						send_flag = true
					}
					if send_flag {
						select {
						//发送数据给连接
						case c.send <- send_msg:
						//关闭连接
						default:
							close(c.send)
							delete(h.connections, c)
						}
					}
				} else {
					select {
					//发送数据给连接
					case c.send <- send_msg:
					//关闭连接
					default:
						close(c.send)
						delete(h.connections, c)
					}
				}

			}
		}
	}
}
