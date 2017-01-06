package middleware

import (
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/session"
	"github.com/lunny/tango"
	"github.com/tango-contrib/binding"
	. "common"

	"fmt"
)

type C struct {
	tango.Ctx
}
type S struct {
	session.Session
}

type B struct {
	binding.Binder
}
type R struct {
	renders.Renderer
}

type BaseHandler struct {
	C
	S						//Session
	B						//binding.Binder
	R 						//renders.Renderer
	Messages []string
	Errors   []string
	Response map[string]interface{}
}

func SetCtx()  {
	
}

func (x *BaseHandler) Handle(ctx *tango.Context) {

	fmt.Println("before")
	ctx.Next()
	fmt.Println("after")
}


func (x *BaseHandler)HTML(name string,T ...map[string]interface{})  {

	sys_params := map[string]interface{}{

		"map_appkey":Cfg.MustValue("map","map_appkey","T6VBZ-YQ7CI-PU5GI-5TMLM-TGK7T-WCBAU"),
		"map_url":Cfg.MustValue("map","map_url","T6VBZ-YQ7CI-PU5GI-5TMLM-TGK7T-WCBAU"),
	}

	x.Render(name,renders.T{
		"C": sys_params, //common
		"P": T,		 //params
	})
}

