// Package screenshot monitors filesystem events for screenshots added to
// directories and emits a screenshot event
package screenshot

// Screenshot describes a screenshot event
type Screenshot struct {
	Description string `json:"description"`
	Filename    string `json:"filename"`
	Hash        string `json:"hash"`
	LocalPath   string `json:"local_path"`
}

// Save saves a processed screenshot in the local database event cache
func (sc Screenshot) Save(dbConn string) error {

	return nil
}
