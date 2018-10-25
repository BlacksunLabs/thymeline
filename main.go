package main

import (
	"fmt"
	"os"

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
		fmt.Printf("cannot access directory %s : %v", dir, err)
		return
	}

	opName, ok, err := ui.GetText("Name of operation")
	if err != nil {
		fmt.Printf("%v", err)
		return
	} else if !ok {
		fmt.Println("user clicked cancel.. Skipping adding operation directory to monitor")
		return
	}
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

	// [BUG]: Not sure why but the Directory Chooser dialog box remains open
	// after following a successful code path and causes the entire program to
	// deadlock.
	//
	// Using os.Exit(0) as a hacky alternative to a typical return in order to
	// mitigate this issue until a solution can be found.
	os.Exit(0)
}

func main() {
	// Handle commandline options
	if newopFlag {
		createop()
	}

	screenshot.WatchDirs()
}
