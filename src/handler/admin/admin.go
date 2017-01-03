package handler




type AdminHandler struct {
	baseHandler
}

func (x *AdminHandler) Get() {

	var params = make(map[string]interface{})
	x.HTML("index.html",params)

}

/*
type RenderAction struct {
	renders.Renderer
}

func (x *RenderAction) Get() {
	x.Render("index.html", renders.T{
		"test": "这个是模板渲染",
	})
}
*/
