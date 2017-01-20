package handler

import (
	. "middleware"
	"fmt"
	. "common"
	"log"

	"gopkg.in/mgo.v2/bson"
	"model"
)

type HomeHandler struct {
	BaseHandler
}

func (x *HomeHandler) Get() {

	params :=map[string]interface{}{
		"a":"a",
		"b":[]byte("ddd"),
	}
	fmt.Print()

	x.HTML("index.html",params)

}
type Person struct {
	Name string
	Phone string
}

func (x *HomeHandler)Post()  {

	mgo := Mgo.DB("test").C("people")
	result := []Person{}

	err := mgo.Insert(&Person{"Ale2", "+55 53 8116 9639"},
		&Person{"Cla2", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}


	err = mgo.Find(bson.M{"name": "Ale2"}).All(&result)

	if err != nil{
		log.Println(err)
	}



	for i,v := range result{
		log.Println(i,":",v.Name,"&&",v.Phone)
	}

	log.Println(result)

	result2 := []model.ChatPositionLog{}
	mgo2 := Mgo.DB("chat").C("chat_position_log")
	mgo2.Find(bson.M{"Uid":1}).All(&result2)
	for i,v := range result2{
		log.Println(i,":",v.Uid,"&&",v.City)
	}
}

func (x *HomeHandler)Say( s interface{} ) string{
	return s.(string)
}