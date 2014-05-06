package gocker

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // used by gorm
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func init() {
	connectionStr := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True",
		GockerCtx.Cfg.Database.User,
		GockerCtx.Cfg.Database.Password,
		GockerCtx.Cfg.Database.Name)
	log.Println("Trying to connect to dababase: ", connectionStr)
	var err error
	GockerCtx.DB, err = gorm.Open("mysql", connectionStr)
	if err != nil {
		log.Panicf("Got error when connect database, the error is '%v'", err)
	}
	createTables()
}

func DBClose() {
	GockerCtx.DB.Close()
}

func DBClear() {
	log.Print("Dropping database tables...")
	defer log.Println("done")
	GockerCtx.DB.DropTable(User{})
	GockerCtx.DB.DropTable(League{})
	GockerCtx.DB.DropTable(Game{})
	GockerCtx.DB.DropTable(GameResult{})
}

// Init default database by dropping recreating tables with default data
func DBDefaultData() {
	DBClear()
	createDefaultData()
}

func createTables() {
	GockerCtx.DB.AutoMigrate(User{})
	GockerCtx.DB.AutoMigrate(League{})
	GockerCtx.DB.AutoMigrate(Game{})
	GockerCtx.DB.AutoMigrate(GameResult{})
}

func createDefaultData() {
	log.Print("Creating default database data...")
	defer log.Print("done")
	createTables()
	//db.Debug().Unscoped().Where("login = ?", "guillaume").Delete(User{})

	user1 := User{
		Login:     "guillaume",
		Password:  "guillaume",
		Name:      "guillaume lazar",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	GockerCtx.DB.Save(&user1)

	user2 := User{
		Login:     "robin",
		Password:  "robin",
		Name:      "robin penea",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	GockerCtx.DB.Save(&user2)

	league := League{
		Name:      "league-test",
		Users:     []User{user1, user2},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	GockerCtx.DB.Save(&league)
}
