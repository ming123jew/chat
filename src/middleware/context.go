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
	tango.Context
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

type Context struct {
	C						//Ctx/context
	S						//Session
	B						//binding.Binder
	R 						//renders.Renderer
	//FormErr  binding.Errors
	Messages []string
	Errors   []string
	Response map[string]interface{}
}


type BaseHandler struct {
	Context
}

func SetCtx()  {
	
}

type HelloHandler struct {}
func (HelloHandler) Handle(ctx *tango.Context) {

	fmt.Println("before")
	ctx.Next()
	fmt.Println("after")
}


func (x *BaseHandler)HTML(name string,T map[string]interface{})  {

	sys_params := map[string]interface{}{

		"map_appkey":Cfg.MustValue("map","map_appkey","T6VBZ-YQ7CI-PU5GI-5TMLM-TGK7T-WCBAU"),
		"map_url":Cfg.MustValue("map","map_url","T6VBZ-YQ7CI-PU5GI-5TMLM-TGK7T-WCBAU"),
	}

	x.Render(name,renders.T{
		"C": sys_params, //common
		"P": T,		 //params
	})
}

