package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"
)

type Walker struct {
	Pack    PackConfig
	TempDir string
}

func (walker *Walker) walk_copy(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("Error accessing path %q: %v\n", path, err)
		return err
	}

	for _, exclude := range walker.Pack.Excludes {
		if strings.Contains(path, exclude) {
			// println("Excluding " + path)

			return nil
		}
	}

	new_path := GetNewPath(path, walker.Pack.Source, walker.TempDir)
	if info.IsDir() {
		os.Mkdir(new_path, 0755)
	} else {
		copy.Copy(path, new_path)
	}

	return nil
}

func GetNewPath(path, from_dir, temp_dir string) string {
	common := path[len(from_dir):] // rest of string without from_dir
	return filepath.Join(temp_dir, common)
}
