package model
import (
	"log"
	. "common"

)
type ChatAppNode struct {

	Id 	int		//`id` int(11) NOT NULL,
	Name 	string		//`name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '节点名称　英文',
	Title 	string		//`title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT ' 对应中文描述',
	Status  int		//`status` tinyint(1) NOT NULL,
	Remark	string	//`remark` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '描述',
	Level	int	//`level` tinyint(2) NOT NULL COMMENT '层级',
	Groupid	int	//`groupid` int(6) NOT NULL,
	Pid	int	//`pid` int(10) NOT NULL,
}

func (self *ChatAppNode)Add(x *ChatAppNode)(int64,error)  {

	res,err := Orm.InsertOne(&ChatAppNode{Id:x.Id,Name:x.Name,Title:x.Title,Status:x.Status,Remark:x.Remark,Level:x.Level,Groupid:x.Groupid,Pid:x.Pid})
	log.Println(res)
	return res,err
}