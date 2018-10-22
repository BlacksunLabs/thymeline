package db

import (
	"go/build"
	"testing"
)

func TestConnect(t *testing.T) {
	_, err := Connect()
	if err != nil {
		t.Errorf("Unable to connect to MongoDB : %v", err)
	}
}

func TestAddToDB(t *testing.T) {
	var projectPath = build.Default.GOPATH + "/src/github.com/BlacksunLabs/thymeline"

	s, err := Connect()
	defer s.Close()

	if err != nil {
		t.Errorf("Unable to connect to MongoDB: %v", err)
	}

	opdir := &OpDir{}
	opdir.Operation.Name = "testdir"
	opdir.Operation.Path = projectPath + "/database/testdir"

	err = opdir.AddToDB(*s)
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveOpDir(t *testing.T) {
	s, err := Connect()
	defer s.Close()

	if err != nil {
		t.Errorf("Unable to connect to MongoDB: %v", err)
	}

	err = RemoveOpDir(*s, "testdir")
	if err != nil {
		t.Errorf("Failed to remove OpDir: %v", err)
	}
}

func TestGetOpDirs(t *testing.T) {
	s, err := Connect()
	defer s.Close()

	if err != nil {
		t.Errorf("Unable to connect to MongoDB: %v", err)
	}

	var dirs []string

	dirs, err = GetOpDirs(*s)
	if err != nil {
		t.Errorf("%v", err)
	}

	t.Log(dirs)
}
