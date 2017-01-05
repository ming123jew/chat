package handler

import (
	. "middleware"

	"fmt"
)

type HomeHandler struct {
	BaseHandler
}

func (x *HomeHandler) Get() {

	params :=map[string]interface{}{
		"a":"a",
	}
	fmt.Print()

	//x.C.ServeJson(params)

	x.HTML("index.html",params)

}