package goker

import (
	"code.google.com/p/gcfg"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var GokerCtx GokerContext

type GokerContext struct {
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
	App struct {
		Host      string
		Port      int
		UrlPrefix string
	}
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.App.Host, c.App.Port)
}

const configFile string = "config.gcfg"

func init() {
	// configure logging system
	log.SetPrefix("goker ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// load config file
	err := gcfg.ReadFileInto(&GokerCtx.Cfg, configFile)
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
