package db

// OpDir hold a map that links directory paths to operation names
type OpDir struct {
	Operations map[string]string `json:"operations"`
}
