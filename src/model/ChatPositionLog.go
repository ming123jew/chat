package model
//use mgo mongodb
import (

	. "common"
	"log"
)

type ChatPositionLog struct {

	//Id 		int		//int(11) NOT NULL,
	Uid 		int		//int(11) NOT NULL,
	Attime 		int64		//int(11) NOT NULL,
	Module		string
	Nation 		string		//varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
	Province 	string		//varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
	City 		string		//varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
	Adcode 		int		//int(6) NULL DEFAULT NULL COMMENT '//行政区ID，六位数字, 前两位是省，中间是市，后面两位是区，比如深圳市ID为440300',
	Addr 		string		//varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
	Lat 		float64		//varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '//火星坐标(gcj02)，腾讯、Google、高德通用',
	Lng 		float64		//varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '//火星坐标(gcj02)，腾讯、Google、高德通用',
	Accuracy 	int		//int(8) NULL DEFAULT NULL COMMENT '误差范围，以米为单位',
	Type 		string
}

func (self *ChatPositionLog)Add(x *ChatPositionLog) (error) {
	mgo := Mgo.DB("chat").C("chat_position_log")
	err := mgo.Insert(&ChatPositionLog{
		x.Uid,
		x.Attime,
		x.Module,
		x.Nation,
		x.Province,
		x.City,
		x.Adcode,
		x.Addr,
		x.Lat,
		x.Lng,
		x.Accuracy,
		x.Type,
	})
	
	if err !=nil{
		log.Println(err)
	}

	return err
}

func (self *ChatPositionLog)Exist() ChatPositionLog {
	mgo := Mgo.DB("chat").C("chat_position_log")
	result := ChatPositionLog{}
	err := mgo.Find(self).One(&result)
	if err !=nil{
		log.Println(err)
	}
	return result
}


