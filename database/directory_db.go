package db

import (
	"fmt"

	bson "github.com/globalsign/mgo/bson"
)

// OpDir hold a map that links directory paths to operation names
type OpDir struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	// Operation map[string]string `json:"operation" bson:"operation"`
	Operation Op
}

// Op is a map of operation names and their local path
type Op struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// AddToDB adds an Operation to the Directory Database
func (od OpDir) AddToDB(s Session) error {
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
func RemoveOpDir(s Session, opname string) error {
	session := s.Copy()
	defer session.Close()

	collection := s.GetCollection(dirDB, dirCollection)

	var ops []OpDir
	err := collection.Find(bson.M{}).All(&ops)
	if err != nil {
		return err
	}

	for _, op := range ops {
		if op.Operation.Name == opname {
			collection.RemoveId(op.ID)
		}
		return nil
	}

	return fmt.Errorf("no operation named %s found", opname)
}

// GetOpDirs queries the OpDir collection for a list of all operation directories
//
// Returns
// A slice of strings containing paths to operation screenshot directories
func GetOpDirs(s Session) ([]string, error) {
	var (
		session    = s.Copy()
		opdirs     []string
		ops        []OpDir
		collection = s.GetCollection(dirDB, dirCollection)
	)
	defer session.Close()

	err := collection.Find(bson.M{}).Select(bson.M{"operation.path": 1}).All(&ops)
	if err != nil {
		return nil, err
	}

	for _, op := range ops {
		opdirs = append(opdirs, op.Operation.Path)
	}

	return opdirs, nil
}

// GetOpNameFromPath gets the Name property of an OpDir and returns if it exists
func GetOpNameFromPath(s Session, path string) (string, error) {
	var (
		session    = s.Copy()
		collection = s.GetCollection(dirDB, dirCollection)
		op         []OpDir
	)
	defer session.Close()

	err := collection.Find(bson.M{}).Select(bson.M{"operation": 1}).All(&op)
	if err != nil {
		return "", err
	}

	for _, op := range op {
		if op.Operation.Path == path {
			return op.Operation.Name, nil
		}
	}

	return "", fmt.Errorf("no operation found with path: %s", path)
}
