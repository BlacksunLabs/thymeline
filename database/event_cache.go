package db

import (
	"time"
	bson "github.com/globalsign/mgo/bson"

)

// Screenshot contains a single screenshot event
type Screenshot struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Timestamp   *time.Time `json:"timestamp"`
	OpName      string     `json:"op_name"`
	Description string     `json:"description"`
	LocalPath   string     `json:"local_path"`
	Hash        string     `json:"hash"`
}

// StoreLocal stores a screenshot to the local Event Cache
func (sc Screenshot) StoreLocal(s *Session) error {
	session := s.Copy()
	defer session.Close()

	collection := s.GetCollection(eventDB, scrotCollection)
	
	err := collection.Insert(sc)
	if err != nil {
		return err
	}
	
	return nil
}