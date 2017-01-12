package middleware

import (
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/session"
	"github.com/lunny/tango"
	"github.com/tango-contrib/binding"
	. "common"

	"fmt"
	"reflect"
	"errors"
	"log"
)
const  (
	SESSION_VALUE_USERLOGIN = "UserLogin" //登录变量
)


type Middler struct {
	Session   *session.Session
	Ctx	  *tango.Context
	Binding
	Renders
	Messages []string
	Errors   []string
	Response map[string]interface{}
}

type MiderlerInterface interface {
	SetMiddler(*session.Session,*tango.Context)

}

func (x *Middler)SetMiddler(s *session.Session,c *tango.Context)  {
	x.Session = s
	x.Ctx = c
	log.Println("heeeee")

}

type Renders struct {
	renders.Renderer
}
type Binding struct {
	binding.Binder
}




//中间件
func MiddleHandler(s *session.Sessions) tango.HandlerFunc  {
	return func(ctx *tango.Context) {
		if action := ctx.Action(); action != nil {
			sess := s.Session(ctx.Req(), ctx.ResponseWriter)

			if miderlerInterface, ok := action.(MiderlerInterface); ok {

				miderlerInterface.SetMiddler(sess,ctx)
			}
		}
		ctx.Next()
	}
}


type BaseHandler struct {
	Middler

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


//判断用户有没有登录
func (x *BaseHandler) IsLogin() bool  {
	s := x.Session.Get(SESSION_VALUE_USERLOGIN)
	if s != nil {
		return true
	}
	return false
}



//判断数组、MAP健值是否存在
func Contains(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("not in")
}

