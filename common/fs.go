package common

import (
	"fmt"
	"os"
	"path/filepath"
)

// ListDir returns a []string of the filepaths of all files in a directory.
func ListDir(path string) ([]string, error) {
	var paths []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		paths = append(paths, path)
		return nil
	})

	if err != nil {
		return []string{}, err
	}

	for _, file := range paths {
		fmt.Println(file)
	}

	return paths, nil
}
