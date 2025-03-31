package internal

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
)

func Pack(pack PackConfig, zip_file string) {
	var walker = to_walker(pack)
	fmt.Printf("\tWalking: %s\n", pack.Source)
	// fmt.Printf("Using TempDir: %s\n", walker.TempDir);

	var err = filepath.Walk(pack.Source, walker.walk_copy)
	if err != nil {
		println("Error walking the path: %v", err)
	}

	fmt.Printf("\tZipping: %s\n", zip_file)
	if err := Zip(walker.TempDir, zip_file); err != nil {
		log.Fatalf("Error zipping: %s", err)
	}

	// fmt.Printf("Deleting TempDir: %s\n", walker.TempDir)
	os.RemoveAll(walker.TempDir)
	println()
}


func ensure_exists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create temp directory base: %w", err)
		}
	}
	return nil
}

