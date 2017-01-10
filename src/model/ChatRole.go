package model
import (
	"log"
	. "common"
	"time"
)
type ChatRole struct {

	Id 		int		//`id` int(6) NOT NULL,
	Name 		string		//`name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
	Pid 		int		//`pid` int(6) NOT NULL,
	Status 		int		//`status` tinyint(1) NOT NULL,
	Remark 		string 		//`remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
	Ename		string		//`ename` varchar(5) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
	Createtime 	int		//`createtime` int(11) NOT NULL,
	Updatetime 	int		//`updatetime` int(11) NOT NULL,


}


func (self *ChatRole)Add(x *ChatRole)(int64,error)  {

	res,err := Orm.InsertOne(&ChatRole{Id:x.Id,Name:x.Name,Pid:x.Pid,Status:x.Status,Remark:x.Remark,Ename:x.Ename,Createtime:int(time.Now().Unix()),Updatetime:0})
	log.Println(res)
	return res,err
}