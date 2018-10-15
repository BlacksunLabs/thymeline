// Package db handles database interaction for thymeline
package db

import (
	"github.com/globalsign/mgo"
)

// Session wraps an pointer to an mgo.Session and adds convenience methods 
type Session struct {
	*mgo.Session
}

const (
	host = "mongodb:27017"
	dirDB = "Directory_DB"
	dirCollection = "OpDirs"
	eventDB = "Event_Cache"
	scrotCollection = "Screenshots"
)

// NewSession returns a pointer to a new Session object given a pointer to an mgo.Session
func NewSession (s *mgo.Session) *Session {
	return &Session{s}
}

// Connect connects to a MongoDB server and initiates a session
//
// Returns
// *Session - pointer to a Session
func Connect() *Session {
	s, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	return NewSession(s)
}

// GetCollection gets a collection from a database given a valid session and returns a pointer to the collection
func (s *Session) GetCollection(db string, collection string) (*mgo.Collection) {
	return s.DB(db).C(collection)
}

