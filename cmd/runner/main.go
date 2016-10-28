package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	for {
		session.DB("test").C("mycollection").Insert(bson.M{"key":"blablabla"})
		time.Sleep(1 * time.Millisecond)
	}
}
