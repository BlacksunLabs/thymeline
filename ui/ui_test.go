package ui

import "testing"

func TestGetOS(t *testing.T) {
	os, err := getOS()
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Log(os)
}

func TestChooseDirectory(t *testing.T) {
	d, err := ChooseDirectory()
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Log(d)
}