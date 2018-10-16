package db

import (
	"fmt"
	bson "github.com/globalsign/mgo/bson"
)

// OpDir hold a map that links directory paths to operation names
type OpDir struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
	Operation map[string]string `json:"operation" bson:"operation"`
}

// AddToDB adds an Operation to the Directory Database
func (od OpDir) AddToDB (s Session) error {
	session := s.Copy()
	defer session.Close()

	collection := session.DB(dirDB).C(dirCollection)
	
	err := collection.Insert(od)
	if err != nil {
		return err
	}

	return nil
}

// RemoveOpDir removes a document from the Database_DB OpDir collection
func RemoveOpDir (s Session, opname string) error {
	session := s.Copy()
	defer session.Close()

	collection := s.GetCollection(dirDB, dirCollection)

	var data []bson.M
	err := collection.Find(bson.M{}).All(&data)
	if err != nil {
		return err
	}

	for _, doc := range data {
		for key, value := range doc {
			if key == "operation" {
				for key, _ := range value.(bson.M) {
					if key == opname {
						id := doc["_id"].(bson.ObjectId)
						collection.RemoveId(id)
					}
					return nil
				}
			}
		}
	}
	return fmt.Errorf("no operation named %s found", opname)
}