package model
import (
	"log"
	. "common"
)
//应用分组
type ChatAppGroup struct {

	Id 		int 		//`id` smallint(3) NOT NULL,
	Name 		string		//`name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '模块名称',
	Title 		string		//`title` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '对应解析',
	Addtime 	int 		//`addtime` int(11) NOT NULL,
	Status 		int 		//`status` tinyint(1) NOT NULL,

}

func (self *ChatAppGroup)Add(x *ChatAppGroup)(int64,error)  {

	res,err := Orm.InsertOne(&ChatAppGroup{Id:x.Id,Name:x.Name,Title:x.Title,Addtime:x.Addtime,Status:x.Status})
	log.Println(res)
	return res,err
}