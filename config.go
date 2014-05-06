package gocker

import (
	"code.google.com/p/gcfg"
	"log"
)

const configFile string = "config.gcfg"

var Cfg Config

type Config struct {
	Database struct {
		Name     string
		User     string
		Password string
	}
}

func init() {
	// configure logging system
	log.SetPrefix("Gocker ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// load config file
	err := gcfg.ReadFileInto(&Cfg, configFile)
	if err != nil {
		log.Fatalf("Failed to parse config file: %s", err)
	}
	log.Println("Config file loaded")
}
