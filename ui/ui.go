// Package ui provides cross-platform UI elements for thymeline
package ui

import (
	"fmt"
	"runtime"
)
const (
	Windows = "windows"
	macOS = "darwin"
	Linux = "linux"
)

func getOS() (string, error) {
	os := runtime.GOOS
	if os == Windows {
		return Windows, nil
	} else if os == macOS {
		return macOS, nil
	} else if os == Linux {
		return Linux, nil
	} else {
		return "", fmt.Errorf("unsupported Operating System: %v", os)
	}
}