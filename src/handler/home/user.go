package handler

import (
	. "middleware"
	"log"
	"fmt"

)


type UserHandler struct {
	BaseHandler
	IsLogin bool


}

type MyStruct struct {
	Id int64
	Name string
}




//默认主页
func (x *UserHandler) Get()  {

	params :=map[string]interface{}{
		"a":"a",
	}

	var mystruct MyStruct
	x.B.Bind(&mystruct)
	log.Println(mystruct.Id,mystruct.Name)

	fmt.Println(x.IP())


	x.HTML("user/index.html",params)
}
