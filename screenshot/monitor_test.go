// +build monitor

package screenshot

import "testing"

func TestWatchDir(t *testing.T) {
	err := WatchDirs()
	if err != nil {
		t.Errorf("%v", err)
	}
}
