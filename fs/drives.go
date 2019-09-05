package fs

import (
	"fmt"
	"io/ioutil"
	"os"
)

// GetDrives returns a list of all drives mounted on a windows machine.
func GetDrivesWindows() ([]string, error) {
	var drives []string
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err != nil {
			return []string{}, err
		}

		drives = append(drives, string(drive))
		err = f.Close()
		if err != nil {
			return []string{}, err
		}
	}
	return drives, nil
}

// GetDrivesDarwin returns a list of all drives mounted on a MacOS machine.
func GetDrivesDarwin() ([]string, error) {
	var drives []string

	mountedDrives, err := ioutil.ReadDir("/Volumes/")
	if err != nil {
		return []string{}, err
	}

	for _, drive := range mountedDrives {
		name := drive.Name()
		if name == "Macintosh HD" {
			continue
		}
		drives = append(drives, drive.Name())
	}

	return drives, nil
}

// GetDrivePath returns the drive path based on the oeprating system.
func GetDrivePath(drive string) string {
	return fmt.Sprintf("/Volumes/%s", drive)
}