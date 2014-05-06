package gocker

import (
    "flag"
	"log"
	"code.google.com/p/gcfg"
	"github.com/jinzhu/gorm"
)

var GockerCtx GockerGockerCtx

type GockerGockerCtx struct {
	Cfg Config
	DB  gorm.DB
}

// CLI parsing data
var cliDbClear = flag.Bool("db-clear", false, "recreate an empty tables structure")
var cliDbDefault = flag.Bool("db-default", false, "clear tables and create default data")

// Configuration flag data
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

func HandleArgs() {
    flag.Parse()
    if *cliDbClear {
        DBClear()
    }
    if *cliDbDefault {
        DBDefaultData()
    }
}

