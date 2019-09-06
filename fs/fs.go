package fs

import (
	"fmt"
	"github.com/xoreo/flash-encrypt/common"
	"os"
	"path/filepath"
	"strings"
)

// ListDir returns a []string of the filepaths of all files in a directory.
func ListDir(rootPath string) ([]string, error) {
	var allPaths []string

	// Walk the directory
	err := filepath.Walk(rootPath, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			allPaths = append(allPaths, path)
		}
		return nil
	})

	if err != nil {
		return []string{}, err
	}

	var paths []string
	for _, path := range allPaths {
		add := true
		for _, prohibitedFile := range common.ProhibitedFiles {
			prohibitedFileFmt := fmt.Sprintf("%s%s", prohibitedFile, common.OSSlash)
			if strings.Contains(path, prohibitedFileFmt) {
				add = false
			}
		}

		if add {
			paths = append(paths, path)
		}
	}

	return paths, nil
}
