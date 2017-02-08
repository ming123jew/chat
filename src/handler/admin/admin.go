package handler

import (
	. "middleware"
	"model"
	"log"
)

const (


)


type Admin struct {
	BaseHandler
}


type AdminLogin struct {
	Admin
	Username string
	Password string
}



func (x *AdminLogin) Get() {
	x.Session.Set("a","okl")
	var params = make(map[string]interface{})
	x.HTML("administrator/login.html",params)
}

func (x *AdminLogin) Post()  {
	var form AdminLogin
	x.Binding.Bind(&form)
	user := model.ChatUser{Username:form.Username,Password:form.Password}

	has, err := user.Exist()
	if err != nil{
		x.Ctx.Write([]byte("账号或密码错误."))
	}

	//验证账号真实性
	if user.Id==0 {
		x.Ctx.Write([]byte("账号不存在."))
		has = false
	}

	//验证权限
	if has == true{
		log.Println(user)
		//x.Test()
	}

}




