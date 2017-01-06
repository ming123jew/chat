package handler

import (
	. "middleware"
	"log"
	"model"
)

const  (
	SESSION_VALUE_USERLOGIN = "UserLogin" //登录变量
)


type UserHandler struct {
	BaseHandler
}

type UserLogin struct {
	UserHandler
	Username string
	Password string

}

func (x *UserHandler) IsLogin() bool  {
	s := x.S.Get(SESSION_VALUE_USERLOGIN)
	if s != nil {
		return true
	}
	return false
}

//默认主页
func (x *UserHandler) Get()  {

	params :=map[string]interface{}{
		"a":"a",
	}

	log.Println(x.IP())

	log.Println(x.S.Get(SESSION_VALUE_USERLOGIN))

	log.Println(x.IsLogin())

	x.HTML("user/index.html",params)
}

func (x *UserLogin) Get(){
	if x.IsLogin() == true{
		x.C.Redirect("/user/index")
	}
	x.HTML("user/login.html")

}

func (x *UserLogin) Post() {
	//username := x.Req().PostForm("")
	//user := &Login{Username:username,Password:password}
	var form UserLogin
	x.B.Bind(&form)

	user := model.ChatUser{Username:form.Username,Password:form.Password}

	has, err := user.Exist()
	if err != nil{
		x.C.Write([]byte("账号或密码错误."))
	}
	if has == true{
		x.S.Set(SESSION_VALUE_USERLOGIN,user)
		log.Println("yes")

	}

	log.Println(user)
	log.Println(form.Username,form.Password)
}

