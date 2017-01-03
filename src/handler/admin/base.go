package handler



import (

"github.com/tango-contrib/renders"
)

type baseHandler struct {
	renders.Renderer
}


func (x *baseHandler)HTML(name string,T map[string]interface{})  {
	x.Render(name, T)
}

