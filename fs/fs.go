package fs

import (
	"fmt"
	"github.com/xoreo/flash-encrypt/common"
	"os"
	"path/filepath"
	"strings"
)

// ListDir returns a []string of the filepaths of all files in a directory.
func ListDir(path string) ([]string, error) {
	var paths []string

	// Walk the directory
	err := filepath.Walk(path, func(filepath string, f os.FileInfo, err error) error {
		for _, prohibitedFile := range common.ProhibitedFiles {
			prohibitedFileFormatted := fmt.Sprintf("%s%s", prohibitedFile, common.OSSlash)
			containsProhibited := strings.Contains(
				filepath,
				prohibitedFileFormatted,
			)
			fmt.Printf("NAME: %s SPRINTF: %s CONTAINS: %t\n\n", prohibitedFileFormatted, filepath, containsProhibited)

			if f.Name() == prohibitedFile || containsProhibited {
				continue
			}

			if !f.IsDir() {
				paths = append(paths, filepath)
			}
		}

		return nil
	})

	if err != nil {
		return []string{}, err
	}

	return paths, nil
}
