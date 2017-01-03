package handler



import (

	"github.com/tango-contrib/renders"
	. "common"

	"fmt"

	"encoding/json"
)

type baseHandler struct {
	renders.Renderer
}

type templateObj struct {

	c map[string]interface{}
}



func (x *baseHandler)HTML(name string,T map[string]interface{})  {
	/**
	cp := make(map[string]map[string]string)
	c := make(map[string]string)
	c["map_url"] = Cfg.MustValue("map","map_url","T6VBZ-YQ7CI-PU5GI-5TMLM-TGK7T-WCBAU")
	cp["map"] = c
	c["map_appkey"] =  Cfg.MustValue("map","map_appkey","T6VBZ-YQ7CI-PU5GI-5TMLM-TGK7T-WCBAU")
	cp["map"] = c

	T["common"] = cp

	fmt.Print(T["common"])
	**/
	var c = new(templateObj)




	fmt.Println(c.c)

	x.Render(name,
		`"map_appkey":`+Cfg.MustValue("map","map_appkey","T6VBZ-YQ7CI-PU5GI-5TMLM-TGK7T-WCBAU"),
		`"map_url":`+Cfg.MustValue("map","map_url","T6VBZ-YQ7CI-PU5GI-5TMLM-TGK7T-WCBAU"),

	)
}

func (x *templateObj)SetCommonParams(key string,value string) ( map[string]interface{} )  {
	var str string
	str += `{"`+key+`": "`+value+`"}`
	var result map[string]interface{}
	if err := json.Unmarshal([]byte( str ), &result); err != nil {

	}
	x.c = result
	return result
}