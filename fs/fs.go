package fs

import (
	"github.com/xoreo/flash-encrypt/common"
	"os"
	"path/filepath"
)

// ListDir returns a []string of the filepaths of all files in a directory.
func ListDir(path string) ([]string, error) {
	var paths []string

	// Walk the directory
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		for _, prohibitedFile := range common.ProhibitedFiles {
			if f.Name() == prohibitedFile {
				continue
			}
		}
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
