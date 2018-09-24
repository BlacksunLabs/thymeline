package ui

import "github.com/martinlindhe/inputbox"

// GetText prompts for user input via a text box given a prompt string
//
// Returns
//
// string entered into textbox
//
// bool indicating whether text input was cancelled (false) or submitted (true)
//
// error if error occurred
func GetText(promptString string) (string, bool, error) {
	got, ok := inputbox.InputBox("Thymeline", promptString, "")
	if !ok {
		return "", false, nil
	}

	return got, true, nil
}
