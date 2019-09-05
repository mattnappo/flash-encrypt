package common

import (
	"os"
	"path/filepath"
)

// ListDir returns a []string of the filepaths of all files in a directory.
func ListDir(path string) ([]string, error) {
	var paths []string

	// Walk the directory
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return []string{}, err
	}

	return paths, nil
}
