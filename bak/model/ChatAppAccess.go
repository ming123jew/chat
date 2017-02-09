package model
import (
	"log"
	. "common"
)
type ChatAppAccess struct {

	Roleid int  		//`roleid` int(10) NOT NULL COMMENT '角色id',
	Nodeid int 		//`nodeid` int(10) NOT NULL COMMENT '节点id',
	Level  int 		//`level` tinyint(1) NOT NULL COMMENT '层级',
	Pid    int 	        //`pid` int(10) NOT NULL COMMENT '上级',
	Module string		//`module` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL

}

func (self *ChatAppAccess)Add(x *ChatAppAccess)(int64,error)  {

	res,err := Orm.InsertOne(&ChatAppAccess{Roleid:x.Roleid,Nodeid:x.Nodeid,Level:x.Level,Pid:x.Pid,Module:x.Module})
	log.Println(res)
	return res,err
}