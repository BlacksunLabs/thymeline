package screenshot

import (
	"fmt"
	"os"

	"github.com/n0ncetonic/thymeline/ui"
)

// GetDescription gets a description string from user input via UI textbox
func GetDescription() (description string, ok bool, err error) {
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
			fmt.Println(description)
			return description, true, nil
		}
	}
	return "", false, fmt.Errorf("giving up after %d retries", retries)
}

// Rename renames a screenshot to a given description
func Rename(oldName string, description string) {

}

// Hash hashes a file with md5
func Hash(f *os.File) (string, error) {

	return "", nil
}
