package handler

type HomeHandler struct {
	baseHandler
}

func (x *HomeHandler) Get() {

	var params = make(map[string]interface{})
	x.HTML("index.html",params)

}