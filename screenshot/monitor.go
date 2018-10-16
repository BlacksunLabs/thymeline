package screenshot

import (
	"time"
	"fmt"
	"github.com/radovskyb/watcher"
)

var w = watcher.New()

// WatchDirs watches directories for new screenshots
func WatchDirs() {
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Create)

	go func() {
		for {
			select {
			case event := <- w.Event:
				fmt.Println(event)
			case err := <-w.Error:
				fmt.Println(err)
			case <-w.Closed:
				return
			}
		}
	}()

	err := w.Start(time.Millisecond * 100)
	if err != nil {
		fmt.Println(err)
	}
}

// AddDir adds a directory to be watched via its path
func AddDir(path string) error {
	err := w.Add(path)
	if err != nil {
		return err
	}
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
