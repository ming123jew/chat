package common

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Orm *xorm.Engine

func SetEngine() *xorm.Engine {
	log.Printf("db initializing...")
	var err error
	var server=Cfg.MustValue("db","server","127.0.0.1")
	var username=Cfg.MustValue("db","username","username")
	var password=Cfg.MustValue("db","password","")
	var db_name=Cfg.MustValue("db","db_name","db_name")
	//var show_sql=Cfg.MustValue("db","show_sql","show_sql")
	var charset=Cfg.MustValue("db","charset","utf8")
	var port=Cfg.MustValue("db","port","3306")

	Orm, err = xorm.NewEngine("mysql", username+":"+password+"@tcp("+server+":"+port+")/"+db_name+"?charset="+charset+"&parseTime=true")
	//Orm, err = xorm.NewEngine("mysql", "root:@/chat?charset=utf8&parseTime=true")
	//fmt.Print(username+":"+password+"@tcp("+server+")/"+db_name+"?charset="+charset+"&parseTime=true")
	if err != nil {
		log.Println(err)
	}
	Orm.TZLocation = time.Local

	return Orm
}


