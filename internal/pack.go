package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Pack(pack PackConfig, pack_dir string) {
	fmt.Printf("Walking: %s\n", pack.Source)

	ensure_exists(pack_dir)
	walker := Walker{
		Pack:    pack,
		TempDir: pack_dir,
	}

	err := filepath.Walk(pack.Source, walker.walk_copy)
	if err != nil {
		log.Fatalf("Error walking the path: %v", err)
	}

	println()
}

func ensure_exists(dir string) {
	os.RemoveAll(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("failed to create temp directory base: %v", err)
		}
	}
}
