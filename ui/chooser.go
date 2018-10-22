package ui

import (
	"github.com/sqweek/dialog"
)

// ChooseDirectory prompts the user to choose a directory via a native GUI prompt
func ChooseDirectory() (string, error) {
	directory, err := dialog.Directory().Title("Thymeline: Choose A Directory").Browse()
	if err != nil {
		return "", err
	}
	return directory, nil
}