package main

import (
	. "middleware"
	. "common"

	"os"
	"github.com/lunny/tango"
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/binding"
	"github.com/tango-contrib/session"
	A "handler/admin"
	H "handler/home"

	"flag"
	"time"
)

func init()  {
	//加载conf
	SetConf()

	//初始化model　orm
	SetEngine()

	//
	SetMgo()

}

//启动session服务
var (
	addr = flag.String("addr", ":8080", "http service address")
	SESS = session.New(session.Options{MaxAge:time.Minute * 20, })

)
func main() {
	//

	//初始化tango
	Tg := tango.Classic()

	//实现核心中间件
	//启动模板服务
	var renders = renders.New(renders.Options{
		Reload: true,
		Directory: "./templates",
	})
	Tg.Use(renders)
	//启动参数到结构体映射
	Tg.Use(binding.Bind())
	Tg.Use(MiddleHandler(SESS))



	//静态文件服务器
	Tg.Use(tango.Static(tango.StaticOptions{RootPath:"static"}))


	//路由
	Tg.Group("/admin", func(g *tango.Group) {
		g.Use( A.AdminHandler )

		g.Route([]string{"GET:Get","POST:Post"},"/login",new(A.AdminLogin))

	})

	Tg.Group("/home", func(g *tango.Group) {
		g.Get("/index", new(H.HomeHandler))
		g.Get("/index2", new(H.HomeHandler).Post)

	})


	Tg.Group("/user", func(g *tango.Group) {
		g.Get("/index", new(H.UserHandler))
		g.Route([]string{"GET:Get","POST:Post"},"/login",new(H.UserLogin))
		g.Route([]string{"GET:Logout"},"/logout",new(H.UserLogin))
		//g.Get("/logout",new(H.UserLogin).Logout)
		//g.Post("/login",new(H.UserLogin))
	})

	go H.ChatHandlerRun()

	//启动websocket
	Tg.Any("/ws",new(H.ChatHandler))

	//启动websocket记录位置
	Tg.Any("/wsposition",new(H.ChatPositionLogHandler))


	//设置访问端口
	os.Setenv("PORT",Cfg.MustValue("common","http_port","8000"))
	os.Setenv("HOST",Cfg.MustValue("common","http_host",""))
	//Tg.RunTLS("F:\\go\\ca\\ca.crt","F:\\go\\ca\\ca.key")
	Tg.Run()


}