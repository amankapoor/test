package common

import (
	"log"

	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
	db      *mgo.Database
)

//Dial creates a mongo session which will be copied/cloned for other operations
func Dial() *mgo.Session {
	log.Println("Connencting to MongoDB...")
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("Unable to Dial to MongoDB on 127.0.0.1:27017: ", err)
	}
	session.SetSafe(&mgo.Safe{})
	session.SetMode(mgo.Monotonic, true)
	log.Println("Connected to testdb...")
	return session
}

//Database return the collection by 'name'
func Database(name string) *mgo.Database {
	return session.DB(name)
}
