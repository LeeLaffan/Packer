package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	packer "packer/internal"
	"time"

	"github.com/BurntSushi/toml"
)

func main() {
	var start_time = time.Now()
	var config_name, pack_name, unpack = parse_args()
	var config, err = read_config(config_name)
	if err != nil { log.Fatalf("ERROR: Could not parse config file: %v\n", err) }

	if unpack {
		return
	} else { 
		if pack_name == "all"{
			for pack_name := range config.Packs { pack(config, pack_name) }
		} else {
			pack(config, pack_name)
		}
	}

	fmt.Printf("Completed in %s\n", time.Since(start_time))
}

func pack(config packer.Config, pack_name string){
	var pack, exists = config.Packs[pack_name]
	if !exists { log.Fatalf("ERROR: Pack '%s' doesn't exist in configuration\n", pack_name) }

	fmt.Printf("Packing: %s\n", pack_name)
	packer.Pack(pack, fmt.Sprintf(config.ZipFmt, pack_name))
}

func parse_args() (string, string, bool) {
	config_name := flag.String("c", "./config.toml", "Config name")
	pack_name := flag.String("n", "", "Pack name")
	unpack := flag.Bool("u", false, "Unpack flag")

	flag.Parse()
	if *pack_name == "" {
		log.Fatalf("ERROR: Argument '-n <pack_name>' is required") 
	}

	fmt.Printf("Args parsed: \n" +
	"\tPack Name : '%v'\n" +
	"\tUnpack    : '%v'\n\n",
	*pack_name, *unpack)

	return *config_name, *pack_name, *unpack
}

func read_config(config_path string) (packer.Config, error) {
	var config packer.Config

	tomlData, err := os.ReadFile(config_path)
	if err != nil { log.Fatalf("ERROR: Failed to read config file: %v", err) }

	if _, err := toml.Decode(string(tomlData), &config); err != nil {
		log.Fatalf("ERROR: Failed to parse TOML: %v", err)
	}

	fmt.Printf("Config parsed: \n" +
	"\tOutputZip : '%v'\n" +
	"\t#Packs    : '%v'\n\n",
	config.ZipFmt, len(config.Packs))

	return config, err
}

