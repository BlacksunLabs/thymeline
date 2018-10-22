package db

import (
	"time"

	bson "github.com/globalsign/mgo/bson"
)

// Screenshot contains a single screenshot event
type ScreenshotEvent struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Timestamp   time.Time     `json:"timestamp"`
	OpName      string        `json:"op_name"`
	Description string        `json:"description"`
	Filename    string        `json:"filename"`
	LocalPath   string        `json:"local_path"`
	Hash        string        `json:"hash"`
}

// SaveScreenshot stores a screenshot to the local Event Cache
func SaveScreenshot(s *Session, screenshot interface{}) error {
	session := s.Copy()
	defer session.Close()

	collection := s.GetCollection(eventDB, scrotCollection)

	err := collection.Insert(screenshot)
	if err != nil {
		return err
	}

	return nil
}
