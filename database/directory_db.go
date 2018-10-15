package db

// OpDir hold a map that links directory paths to operation names
type OpDir struct {
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