package screenshot

import (
	"fmt"
	"time"

	"github.com/BlacksunLabs/thymeline/database"
	"github.com/radovskyb/watcher"
)

var w = watcher.New()

// WatchDirs watches directories for new screenshots
func WatchDirs() error {
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Create)

	go func() {
		for {
			select {
			case event := <-w.Event:
				ok, err := Process(event.Path)
				if !ok || err != nil {
					fmt.Printf("%v", err)
				}
			case err := <-w.Error:
				fmt.Println(err)
			case <-w.Closed:
				return
			}
		}
	}()

	opdirs, err := getOpDirs()
	if err != nil {
		return err
	}

	for _, dir := range opdirs {
		err = AddDir(w, dir)
		if err != nil {
			return fmt.Errorf("Unable to add %s : %v", dir, err)
		}
	}

	err = w.Start(time.Millisecond * 100)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// AddDir adds a directory to be watched via its path
func AddDir(w *watcher.Watcher, path string) error {
	err := w.Add(path)
	if err != nil {
		return err
	}
	return nil
}

// GetOpDirs gets all current OpDirs from Directory_DB
func getOpDirs() ([]string, error) {
	s, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer s.Close()

	opdirs, err := db.GetOpDirs(*s)
	if err != nil {
		return nil, err
	}

	return opdirs, nil
}

// IsOpDir determines whether a directory is part of an active operation
func IsOpDir(path string) (bool, error) {
	return true, nil
}
