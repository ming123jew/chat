package main

import (
	. "common"
	"os"
	"github.com/lunny/tango"
	"github.com/tango-contrib/renders"
	A "handler/admin"
	H "handler/home"
)

func init()  {
	//加载conf
	SetConf()

	//初始化model　orm
	SetEngine()

}


func main() {

	//初始化tango
	Tg := tango.Classic()

	//静态文件服务器
	Tg.Use(tango.Static(tango.StaticOptions{Prefix:"static"}))

	//启动模板服务
	Tg.Use(renders.New(renders.Options{
		Reload: true,
		Directory: "./templates/home",
	}))

	//路由
	Tg.Group("/admin", func(g *tango.Group) {
		g.Get("/index", new(A.AdminHandler))
		g.Get("/1", func() string{

			return "/1"
		})
	})

	Tg.Group("/home", func(g *tango.Group) {
		g.Get("/index", new(H.HomeHandler))

	})
	//设置访问端口
	os.Setenv("PORT",Cfg.MustValue("common","http_port","8000"))
	os.Setenv("HOST",Cfg.MustValue("common","http_host","127.0.0.1"))
	Tg.Run()
}