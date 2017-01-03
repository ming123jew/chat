package middleware

import (
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/session"
	"net/http"
	"github.com/lunny/tango"
)

type Context struct {
	renders.Renderer
	Tg       *tango.Tango
	C        tango.Context
	S        session.Session
	R        *http.Request
	W        http.ResponseWriter
	//FormErr  binding.Errors
	Messages []string
	Errors   []string
	Response map[string]interface{}
}

