package handler



import (

	"github.com/tango-contrib/renders"
	. "common"
)

type baseHandler struct {
	renders.Renderer
}



func (x *baseHandler)HTML(name string,T map[string]interface{})  {

	sys_params := map[string]string{

		"map_appkey":Cfg.MustValue("map","map_appkey","T6VBZ-YQ7CI-PU5GI-5TMLM-TGK7T-WCBAU"),
		"map_url":Cfg.MustValue("map","map_url","T6VBZ-YQ7CI-PU5GI-5TMLM-TGK7T-WCBAU"),
	}

	x.Render(name,renders.T{
		"C": sys_params,
		"P": T,
	})
}
