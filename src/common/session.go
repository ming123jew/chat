package common

import (
	"github.com/tango-contrib/session"
	"github.com/lunny/tango"
	"time"

	"fmt"
)

type Sess struct {
	Session   *session.Session
}

type SessInterface interface {
	SetMiddler(*session.Session)
}

func (x *Sess)SetMiddler(s *session.Session)  {
	x.Session = s
}

func NewSess() *Sess {

	return &Sess{}
}

func Se()  *Sess {
	return &Sess{}
}

//中间件
func (x *Sess) Handle(ctx *tango.Context) {
	s := session.New(session.Options{MaxAge:time.Minute * 20, })
	if action := ctx.Action(); action != nil {
		sess := s.Session(ctx.Req(), ctx.ResponseWriter)
		if miderlerInterface, ok := action.(SessInterface); ok {
			miderlerInterface.SetMiddler(sess)
		}
	}
	fmt.Println("nesssst")
	ctx.Next()
}