package db

import "time"

// Screenshot contains a single screenshot event
type Screenshot struct {
	Timestamp   *time.Time `json:"timestamp"`
	OpName      string     `json:"op_name"`
	Description string     `json:"description"`
	LocalPath   string     `json:"local_path"`
	Hash        string     `json:"hash"`
}
