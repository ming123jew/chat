package main

import (
	. "common"
	"github.com/lunny/tango"
	//"fmt"

	"os"
	//"fmt"
	"model"
	"fmt"
)

func init()  {
	//加载conf
	SetConf()

	//初始化model　orm
	SetEngine()
}

type EventAction struct {
	tango.Json
}

func (EventAction)Get() interface{}  {

	return map[string]string{"a":"b"}
}


func main() {

	user :=  &model.ChatUser{Id:1}

	user,_ = user.GetChatUserById(1)
	fmt.Print(user)

	t := tango.Classic()
	t.Get("/", new(EventAction))
	//t.Use(MyHandler())

	//静态文件服务器
	t.Use(tango.Static(tango.StaticOptions{Prefix:"static"}))


	//设置访问端口
	os.Setenv("PORT",Cfg.MustValue("common","http_port","8000"))
	os.Setenv("HOST",Cfg.MustValue("common","http_host","127.0.0.1"))
	t.Run()
}