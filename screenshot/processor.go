package screenshot

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/BlacksunLabs/thymeline/database"
)

// Process processes a screenshot event and stores the event in the local database
func Process(path string) (ok bool, err error) {
	s, err := db.Connect()
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer s.Close()

	sc := &Screenshot{}

	for i := 0; i < 3; i++ {
		ok, err = sc.GetDescription()
		if err != nil {
			return false, err
		} else if !ok {
			if i == 2 {
				return false, fmt.Errorf("user cancelled 3 times")
			}
			continue
		} else {
			break
		}
	}

	sc.Rename(path)

	err = sc.HashFile()
	if err != nil {
		return false, err
	}

	opname, err := db.GetOpNameFromPath(*s, filepath.Dir(path))
	if err != nil {
		fmt.Printf("%v", err)
	}
	sc.OpName = opname

	sc.Timestamp = time.Now().Local()
	err = db.SaveScreenshot(s, sc)
	if err != nil {
		return false, err
	}

	return true, nil
}
