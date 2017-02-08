package main

import (
	"log"
	//"fmt"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"reflect"
)

type Person struct {
	Name string
	Phone string
}

type base struct {
	a interface{}
}

func main() {
	b := base{}
	r := make([]int,1)
	b.a = &r
	targetValue := reflect.ValueOf(b.a)
	log.Println( "he",targetValue.Len() )
/*
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)*/
}