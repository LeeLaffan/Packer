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

func to_walker(pack PackConfig) Walker {
	temp_dir, _ := os.MkdirTemp(os.TempDir(), "")
	err := ensure_exists(temp_dir) 
	if err != nil { panic(err) }

	return Walker {
		Pack: pack,
		TempDir: temp_dir,
	}
}

func (walker *Walker) walk_copy(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("Error accessing path %q: %v\n", path, err)
		return err
	}

	for _, exclude := range walker.Pack.Excludes {
		if strings.Contains(path, exclude) {
			return nil
		}
	}

	var new_path = GetNewPath(path, walker.Pack.Source, walker.TempDir)
	if info.IsDir() {
		os.Mkdir(new_path, 0755)
	} else {
		copy.Copy(path, new_path)
	}

	return nil
}

func GetNewPath(path string, from_dir string, temp_dir string) string {
	var common = path[len(from_dir):] // rest of string without from_dir
	return filepath.Join(temp_dir, common)
}
