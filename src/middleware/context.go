package middleware

import (
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/session"
	"github.com/lunny/tango"
	"github.com/tango-contrib/binding"
	. "common"
	"reflect"
	"errors"
	"log"
	"strings"
)
const  (
	SESSION_VALUE_USERLOGIN = "UserLogin" 	//登录session变量

	SESSION_VALUE_ADMINLOGIN = "AdminLogin" //后台登录session变量
	SESSION_VALUE_ADMINLOGIN_ROLE = "AdminLoginRoleid"	//后台登录roleid session变量
	ADMIN_FLAG = "admin"			//后台入口标示
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
	setMiddler(*session.Session,*tango.Context)
	setRbac()

}

func (x *Middler)setMiddler(s *session.Session,c *tango.Context)  {
	x.Session = s
	x.Ctx = c
}

func (x *Middler)setRbac()  {
	//暂未作使用,预留
	//判断是否是后台入口如是则进行检测权限
	if USER_AUTH_TYPE > 0 {

		params := strings.Split(strings.ToLower(strings.Split(x.Ctx.Req().RequestURI, "?")[0]), "/")
		//后台验证
		if params[1]==ADMIN_FLAG {
			if params[2] != "login" {
				//判断是否已经登陆
				if !x.IsLoginAdmin() {
					x.Ctx.Redirect("login")
					return
				}
			}
			log.Println(params)
			if  CheckAccess(params) {

				if USER_AUTH_TYPE==1 {
					//session登录认证

				}else if USER_AUTH_TYPE==2{
					//实时认证

					roleid := x.GetAdminSesionRoleid()
					log.Println(roleid)
				}


			}
		}

	}

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

				miderlerInterface.setMiddler(sess,ctx)
				miderlerInterface.setRbac()
			}
		}
		ctx.Next()
	}
}


func (x *Middler)HTML(name string,T ...map[string]interface{})  {

	sys_params := map[string]interface{}{

		"map_appkey":Cfg.MustValue("map","map_appkey",""),
		"map_url":Cfg.MustValue("map","map_url",""),
		"static_url":Cfg.MustValue("common","static_url",""),

	}
	T2 := make(map[string]interface{})
	for _,v := range T{
		T2 = v
	}

	x.Render(name,renders.T{
		"C": sys_params, //common
		"P": T2,		 //params
	})
}


//判断用户有没有登录
func (x *Middler) IsLogin() bool  {
	s := x.Session.Get(SESSION_VALUE_USERLOGIN)
	if s != nil {
		return true
	}
	return false
}

func(x *Middler) GetUserSessionInfo() interface{}{

	s := x.Session.Get(SESSION_VALUE_USERLOGIN)
	return s
}

func (x *Middler) IsLoginAdmin() bool {
	s := x.Session.Get(SESSION_VALUE_ADMINLOGIN)
	if s != nil {
		return true
	}
	return false
}

func(x *Middler) GetAdminSessionInfo() interface{}{

	s := x.Session.Get(SESSION_VALUE_ADMINLOGIN)
	return s
}

func (x *Middler)GetAdminSesionRoleid() interface{} {
	s := x.Session.Get(SESSION_VALUE_ADMINLOGIN_ROLE)
	return s
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



type BaseHandler struct {
	Middler

}