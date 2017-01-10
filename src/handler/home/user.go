package handler

import (
	. "middleware"
	"log"
	"model"
)



type UserHandler struct {
	BaseHandler
}

type UserLogin struct {
	UserHandler
	Username string
	Password string

}

func init()  {

}



//默认主页
func (x *UserHandler) Get()  {

	//access := &model.ChatRoleUser{1,1}
	//access.Add(access)
	new(Rb).GetAccessList(1)


	log.Println(x.IP())

	log.Println(x.S.Get(SESSION_VALUE_USERLOGIN))

	log.Println(x.IsLogin())

	var user = x.S.Get(SESSION_VALUE_USERLOGIN)

	params :=map[string]interface{}{
		"a"	:"a",
		"user"	:user,
	}
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
		x.C.Redirect("/user/index")
	}

	log.Println(user)
	log.Println(form.Username,form.Password)
}

