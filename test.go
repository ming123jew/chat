package main

import (
	"strings"
	"log"
	"reflect"
	"errors"
	"github.com/tango-contrib/session"
	"time"
	. "middle"
	"github.com/lunny/tango"
)

const   (
	REQUIRE_AUTH_MODULE = "Public,Node"
	NOT_AUTH_MODULE =""
)

type   S  struct {
	session.Session
}



func (x *S)Handle()  {
	session.New(session.Options{
		MaxAge:time.Minute * 20,
	})
	x.Set("a","b")
}

var (
	Sess = session.Default()
)
func main()  {


	//初始化tango
	Tg := tango.Classic()

	Tg.Use( Hand(Sess) )

	Tg.Get("/user2",new(User),)

	Tg.Run()


	var module map[string][]string
	//var action map[string][]string
	module = make(map[string][]string)
	if REQUIRE_AUTH_MODULE!="" {
		module["yes"] = strings.Split(REQUIRE_AUTH_MODULE,",")
	}else{
		module["no"] = strings.Split(NOT_AUTH_MODULE,",")
	}

	if ok,_ := Contains("Publicd",module["yes"]) ;ok{
		log.Println("yes")
	}
	if ok,_ := Contains("Publicd",module["no"]) ;ok{
		log.Println("yes")
	}
	if module["no"]==nil{
		log.Println("no")
	}

}

func Contains(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("not in")
}
