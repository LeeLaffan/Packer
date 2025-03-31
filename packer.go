package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	packer "packer/internal"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/mholt/archiver"
)

type Args struct {
	config_name string
	pack_name   string
	unpack      bool
	zip         bool
	del         bool
}

func main() {
	start_time := time.Now()
	args := parse_args()
	config := read_config(args.config_name)

	if args.unpack {
		return
	} else {
		if args.pack_name == "all" {
			for all_name := range config.Packs {
				args.pack_name = all_name
				process_pack(config, args)
			}
		} else {
			process_pack(config, args)
		}
	}

	fmt.Printf("Completed in %s\n", time.Since(start_time))
}

func process_pack(config packer.Config, args Args) {
	pack, exists := config.Packs[args.pack_name]
	if !exists {
		log.Fatalf("ERROR: Pack '%s' doesn't exist in configuration", args.pack_name)
	}

	pack_dir := fmt.Sprintf(config.PackDir, args.pack_name)

	fmt.Printf("Packing: %s\n", args.pack_name)
	packer.Pack(pack, pack_dir)

	if args.zip {
		zip_file := fmt.Sprintf(config.ZipFmt, args.pack_name)
		fmt.Printf("Zipping: %s\n", zip_file)

		os.RemoveAll(zip_file)
		if err := archiver.Archive([]string{pack_dir}, zip_file); err != nil {
			log.Fatalf("Error zipping: %s", err)
		}
		if args.del {
			os.RemoveAll(pack_dir)
		}
	}
}

func parse_args() Args {
	config_name := flag.String("c", "./config.toml", "Config name")
	pack_name := flag.String("n", "", "Pack name")
	unpack := flag.Bool("u", false, "Unpack flag")
	zip := flag.Bool("z", false, "Zip flag")
	del := flag.Bool("d", false, "Delete after zip flag")

	flag.Parse()
	if *pack_name == "" {
		log.Fatalf("ERROR: Argument '-n <pack_name>' is required")
	}

	args := Args{
		config_name: *config_name,
		pack_name:   *pack_name,
		unpack:      *unpack,
		zip:         *zip,
		del:         *del,
	}
	fmt.Printf("Parsed Args: %+v\n", args)
	return args
}

func read_config(config_path string) packer.Config {
	var config packer.Config

	tomlData, err := os.ReadFile(config_path)
	if err != nil {
		log.Fatalf("ERROR: Failed to read config file: %v", err)
	}

	if _, err := toml.Decode(string(tomlData), &config); err != nil {
		log.Fatalf("ERROR: Failed to parse TOML: %v", err)
	}

	if config.PackDir == "" {
		temp_dir, _ := os.MkdirTemp(os.TempDir(), "")
		temp_dir = filepath.Join(temp_dir, "%s")
		fmt.Printf("PackDir is empty, using temp dir: %s\n", temp_dir)
		config.PackDir = temp_dir
	}

	fmt.Printf("Config parsed: %+v\n", config)
	return config
}
