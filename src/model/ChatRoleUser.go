package model
import (
	"log"
	. "common"
)
type ChatRoleUser struct {

	Roleid int 	//`roleid` int(11) NOT NULL,
	Userid int 	//`userid` int(11) NOT NULL

}

func (self *ChatRoleUser)Add(x *ChatRoleUser)(int64,error)  {

	res,err := Orm.InsertOne(&ChatRoleUser{Roleid:x.Roleid,Userid:x.Userid})
	log.Println(res)
	return res,err
}