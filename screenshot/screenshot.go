// Package screenshot monitors filesystem events for screenshots added to
// directories and emits a screenshot event
package screenshot

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/n0ncetonic/thymeline/ui"
)

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

// GetDescription gets a description string from user input via UI textbox
func (sc *Screenshot) GetDescription() (description string, ok bool, err error) {
	retries := 3

	for i := 0; i < retries; i++ {
		description, ok, err := ui.GetText("Enter a description")
		if err != nil {
			fmt.Printf("%v", err)
		} else if !ok {
			// fmt.Println("User clicked cancel")
			return "", false, nil
		} else if len(description) < 1 {
			// fmt.Println("Empty input")
		} else {
			sc.Description = description
			return description, true, nil
		}
	}
	return "", false, fmt.Errorf("giving up after %d retries", retries)
}

// HashFile hashes a file with md5 and adds the hash to a Screenshot
func (sc *Screenshot) HashFile() error {
	f, err := os.Open(sc.LocalPath)
	if err != nil {
		return err
	}
	defer f.Close()

	h := md5.New()

	if _, err := io.Copy(h, f); err != nil {
		return err
	}

	hashInBytes := h.Sum(nil)[:16]
	sc.Hash = hex.EncodeToString(hashInBytes)
	return nil
}

// Rename renames a screenshot to a given description
func (sc *Screenshot) Rename(oldName string, description string) {
	dir := filepath.Dir(oldName)

	os.Rename(oldName, dir+description+".png")
	sc.Filename = description + ".png"
	sc.LocalPath = dir + sc.Filename
}
