// app/database/setup.go

package database

import "gopkg.in/mgo.v2"

/*
Database session
*/
var Session *mgo.Session

/*
Models connection
*/
var Users *mgo.Collection
var Speech *mgo.Collection

/*
Init database
*/
func Init(uri, dbname string) error {
	session, err := mgo.Dial(uri)
	if err != nil {
		return err
	}

	// See https://godoc.org/labix.org/v2/mgo#Session.SetMode
	session.SetMode(mgo.Monotonic, true)

	// Expose session and models
	Session = session
	Users = session.DB(dbname).C("users")
	Speech = session.DB(dbname).C("speeches")

	return nil
}