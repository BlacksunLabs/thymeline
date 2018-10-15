package screenshot

import (
	"fmt"
)

// Process processes a screenshot event and stores the event in the local database
func Process(path string) error {
	sc := &Screenshot{}
	description, ok, err := sc.GetDescription()
	if err != nil {
		return err
	} else if !ok {
		return fmt.Errorf("user clicked cancel button")
	}

	sc.Rename(path, description)

	err = sc.HashFile()
	if err != nil {
		return err
	}

	return nil
}
