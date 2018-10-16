package db

import "testing"

func TestConnect(t *testing.T) {
	_, err := Connect()
	if err != nil {
		t.Errorf("Unable to connect to MongoDB : %v", err)
	}
}

func TestAddToDB(t *testing.T) {
	s, err := Connect()
	if err != nil {
		t.Errorf("Unable to connect to MongoDB: %v", err)
	}

	opdir := &OpDir{}
	testDir := make(map[string]string)
	testDir["testdir"] = "/Path/to/non-existent/directory"
	opdir.Operation = testDir
	opdir.AddToDB(*s)
}