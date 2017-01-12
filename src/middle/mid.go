package middle

import (
	"github.com/tango-contrib/session"
	"github.com/lunny/tango"
)

type Mid struct {
	Session *session.Session
	Ctx     *tango.Context
}

type Midinterface interface {
	SetMid(*session.Session,*tango.Context)
}

func (x *Mid)SetMid(s *session.Session,c *tango.Context)  {
	x.Session = s
	x.Ctx = c
}

func Hand(s *session.Sessions)tango.HandlerFunc  {
	return func(ctx *tango.Context) {
		if action := ctx.Action(); action != nil {
			sess := s.Session(ctx.Req(), ctx.ResponseWriter)

			if miderlerInterface, ok := action.(Midinterface); ok {

				miderlerInterface.SetMid(sess,ctx)
			}
		}
		ctx.Next()
	}
}