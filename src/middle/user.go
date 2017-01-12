package middle

import "log"

type User struct {
	Mid
	Rbac
}

func (x *User)Get()  {
	x.Check()
	x.Session.Set("s","dddddd")
	log.Println("set ok")
}
