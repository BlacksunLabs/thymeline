package ui

import "testing"

func TestGetOS(t *testing.T) {
	os, err := getOS()
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Log(os)
}