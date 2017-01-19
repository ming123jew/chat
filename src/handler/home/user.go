package handler

import (
	. "middleware"
	"log"
	"model"
	"time"
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
time.Sleep(time.Second*3)
	//x.Ctx.Write([]byte("{status:1,data:123}"))

	//b := map[string]string{
	//	"status":"1",
	//}
	c := []map[string]string{
		0 :   {
			"title":"发中国雾霾财，这家美国公司，靠一个口罩赚300亿",
			"thumb":"/images/post-img1.jpg",
			"url":"",
			"viewnum":"2",
			"desc":"发中国雾霾财，这家美国公司，靠一个口罩赚300亿",
		},
		1 :   {
			"title":"论用10年房价、收入涨幅计算心理阴影面积",
			"thumb":"/images/post-img2.png",
			"url":"",
			"viewnum":"20",
			"desc":"夕阳下，是房价飞奔在工薪族心理留下的巨大阴影",
		},
	}
	//a["res"] = b
	a := map[string]interface{}{
		"list":c,
		"status":200,
	}
	x.Ctx.Header()
	x.Ctx.ServeJson(a)
	//access := &model.ChatRoleUser{1,1}
	//access.Add(access)

	//x.Rbac.AccessDecision( "User","Get" )
	//x.GetModuleAccessList(1,"test")
/*
	log.Println("method",x.Ctx.Req().Method )


	log.Println(x.Ctx.IP())

	log.Println(x.Session.Get(SESSION_VALUE_USERLOGIN))

	log.Println(x.IsLogin())

	var user = x.Session.Get(SESSION_VALUE_USERLOGIN)

	params :=map[string]interface{}{
		"a"	:"a",
		"user"	:user,
	}

	user2 := model.ChatUser{}
	c,_ := user2.Count()
	log.Println("user count:",c)
	x.HTML("user/index.html",params)*/
}

func (x *UserLogin) Get(){
	if x.IsLogin() == true{
		x.Ctx.Redirect("/user/index")
	}
	a := make(map[string]interface{})
	x.HTML("user/login.html",a)

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

