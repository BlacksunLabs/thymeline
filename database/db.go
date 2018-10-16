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
	host = "localhost:27017"
	dirDB = "Directory_DB"
	dirCollection = "OpDirs"
	eventDB = "Event_Cache"
	scrotCollection = "Screenshots"
)

// newSession returns a pointer to a new Session object given a pointer to an mgo.Session
func newSession (s *mgo.Session) *Session {
	return &Session{s}
}

// Connect connects to a MongoDB server and initiates a session
//
// Returns
// *Session - pointer to a Session
// error if any
func Connect() (*Session, error) {
	s, err := mgo.Dial(host)
	if err != nil {
		return &Session{}, err
	}

	return newSession(s), nil
}

// GetCollection gets a collection from a database given a valid session and returns a pointer to the collection
func (s *Session) GetCollection(db string, collection string) (*mgo.Collection) {
	return s.DB(db).C(collection)
}

