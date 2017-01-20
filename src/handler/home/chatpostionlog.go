package handler

import (
	. "middleware"
	"log"
	"model"
	"time"
	"encoding/json"
)

type ChatPositionLogHandler struct {

	Huber
	Middler
}


func (x *ChatPositionLogHandler)Get()  {

	ws,err := Upgrader.Upgrade(x.Ctx.ResponseWriter,x.Ctx.Req(),nil)
	if err != nil {
		log.Println("error:",err)
		return
	}

	//初始化Connection
	c := &Connection{Send: make(chan []byte, 256), Ws: ws, Auth: false,}
	defer func() {
		//h.unregister <- c
		log.Println("conn2 close")
		c.Ws.Close()
	}()
	c.Ws.SetReadLimit(MaxMessageSize)
	//c.Ws.SetReadDeadline(time.Now().Add(PongWait))
	c.Ws.SetPongHandler(func(string) error { c.Ws.SetReadDeadline(time.Now().Add(PongWait)); return nil })
	for {
		_, message, err := c.Ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		var obj model.ChatPositionLog
		obj = model.ChatPositionLog{Uid:1,Attime:time.Now().Unix()}
		json.Unmarshal(message,&obj)
		chatpositionlog := new(model.ChatPositionLog)
		chatpositionlog.Add(&obj)
	}


}