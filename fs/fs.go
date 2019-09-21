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

	// Don't add prohibited files (defined in common)
	for _, path := range allPaths {
		add := true

		// For each prohibited file, don't add it.
		for _, prohibitedFile := range common.ProhibitedFiles {
			prohibitedFileFmt := fmt.Sprintf("%s%s",	 prohibitedFile, common.OSSlash)

			tempParts := strings.Split(path, common.OSSlash)
			firstTwoChars := tempParts[len(tempParts) - 1][:2]

			if strings.Contains(path, prohibitedFileFmt) || firstTwoChars == "._" {
				add = false
			}
		}

		// Add the file
		if add {
			paths = append(paths, path)
		}
	}

	return paths, nil
}
