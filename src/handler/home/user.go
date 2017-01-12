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



//默认主页
func (x *UserHandler) Get()  {

	//access := &model.ChatRoleUser{1,1}
	//access.Add(access)

	//x.Rbac.AccessDecision( "User","Get" )
	//x.GetModuleAccessList(1,"test")

	log.Println("method",x.Ctx.Req().Method )


	log.Println(x.Ctx.IP())

	log.Println(x.Session.Get(SESSION_VALUE_USERLOGIN))

	log.Println(x.IsLogin())

	var user = x.Session.Get(SESSION_VALUE_USERLOGIN)

	params :=map[string]interface{}{
		"a"	:"a",
		"user"	:user,
	}
	x.HTML("user/index.html",params)
}

func (x *UserLogin) Get(){
	if x.IsLogin() == true{
		x.Ctx.Redirect("/user/index")
	}
	x.HTML("user/login.html")

}


//登录操作
func (x *UserLogin) Post() {
	//username := x.Req().PostForm("")
	//user := &Login{Username:username,Password:password}
	var form UserLogin
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

	if has == true{

		x.Session.Set(SESSION_VALUE_USERLOGIN,user)
		log.Println("yes")
		x.Ctx.Redirect("/home/index")
	}

	log.Println(user)
	log.Println(form.Username,form.Password)
}
//登出操作

func (x *UserLogin) Logout() {
	x.Session.Set(SESSION_VALUE_USERLOGIN,nil)
	x.Ctx.Redirect("/user/login")
}

