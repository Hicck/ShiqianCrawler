package utils

import (
	"gopkg.in/mgo.v2"
	"log"
)

var Session *mgo.Session

func init(){
	var err error
	Session, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	Session.SetMode(mgo.Monotonic, true)
	log.Printf("################ init mongodb scuess ##############")
}

