package gocker

import (
	"code.google.com/p/gcfg"
	"github.com/jinzhu/gorm"
	"log"
)

var GockerCtx GockerGockerCtx

type GockerGockerCtx struct {
	Cfg Config
	DB  gorm.DB
}

type Config struct {
	Database struct {
		Name     string
		User     string
		Password string
	}
}

const configFile string = "config.gcfg"

func init() {
	// configure logging system
	log.SetPrefix("Gocker ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// load config file
	err := gcfg.ReadFileInto(&GockerCtx.Cfg, configFile)
	if err != nil {
		log.Fatalf("Failed to parse config file: %s", err)
	}
	log.Println("Config file loaded")
}
