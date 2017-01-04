package main

import (
	. "common"
	. "middleware"
	"os"
	"github.com/lunny/tango"
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/binding"
	A "handler/admin"
	H "handler/home"
)

func init()  {
	//加载conf
	SetConf()

	//初始化model　orm
	SetEngine()

	//初始化ctx
	SetCtx()
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

	//启动参数到结构体映射
	Tg.Use(binding.Bind())

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


	Tg.Group("/user", func(g *tango.Group) {
		g.Get("/index", new(H.UserHandler))

	})

	//Tg.Use(new(HelloHandler))

	//设置访问端口
	os.Setenv("PORT",Cfg.MustValue("common","http_port","8000"))
	os.Setenv("HOST",Cfg.MustValue("common","http_host",""))
	Tg.Run()
}