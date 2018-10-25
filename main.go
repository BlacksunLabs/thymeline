package main

import (
	"fmt"

	"github.com/BlacksunLabs/thymeline/database"
	"github.com/BlacksunLabs/thymeline/screenshot"
	"github.com/BlacksunLabs/thymeline/ui"
	flag "github.com/spf13/pflag"
)

// Commandline flags
var (
	newopFlag bool
)

// Parse commandline options
func init() {
	flag.BoolVarP(&newopFlag, "new", "n", false, "used to create and monitor a new operation directory")
	flag.Parse()
}

func createop() {
	dir, err := ui.ChooseDirectory()
	if err != nil {
		fmt.Printf("Unable to create directory `%s` : %v", dir, err)
		return
	}

	opName, ok, err := ui.GetText("Name of operation")
	if err != nil {
		fmt.Printf("%v", err)
		return
	} else if !ok {
		fmt.Println("user clicked cancel.. Skipping adding operation directory to monitor")
	} else { // No error, user didn't cancel. Try adding the op
		s, err := db.Connect()
		if err != nil {
			fmt.Printf("unable to connect to database: %v", err)
			return
		}
		defer s.Close()

		opdir := db.OpDir{
			Operation: db.Op{
				Name: opName,
				Path: dir,
			},
		}
		err = opdir.AddToDB(*s)
		if err != nil {
			fmt.Printf("unable to add operation to database: %v", err)
		}
		return
	}
}

func main() {
	// Handle commandline options
	if newopFlag {
		createop()
	}

	screenshot.WatchDirs()
}
