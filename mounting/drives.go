package mounting

import "os"

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