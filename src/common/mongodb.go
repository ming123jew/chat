package common

import (
	"gopkg.in/mgo.v2"
)

var Mgo *mgo.Session


func SetMgo() {

	m, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	//defer Mgo.Close()
	// Optional. Switch the session to a monotonic behavior.
	m.SetMode(mgo.Monotonic, true)

	Mgo = m

}