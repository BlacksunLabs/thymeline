package screenshot

// WatchDir watches a directory for new screenshots
func WatchDir(path string) error {
	return nil
}

// IsOpDir determines whether a directory is part of an active operation
func IsOpDir(path string) (bool, error) {
	return true, nil
}

// GetOpNameFromDir gets an operation name for a given directory
func GetOpNameFromDir(path string) (string, error) {
	return "", nil
}
