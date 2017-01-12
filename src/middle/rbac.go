package middle

import "log"

type Rbac struct {
	Mid
}

func (x *Rbac)Check()  {
	s := x.Session.Get("s")
	log.Println("session:s",s)
}