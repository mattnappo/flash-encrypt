package fs

import (
	"testing"
)

//func TestGetDrivesWindows(t *testing.T) {
//	drives, err := GetDrivesWindows()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	for i, drive := range drives {
//		t.Logf("[%d] %s\n", i, drive)
//	}
}

func TestGetDrivesDarwin(t *testing.T) {
	drives, err := GetDrivesDarwin()
	if err != nil {
		t.Fatal(err)
	}

	for i, drive := range drives {
		t.Logf("[%d] %s\n", i, drive)
	}
}
