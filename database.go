package gocker

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // used by gorm
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var DB gorm.DB

func init() {
	connectionStr := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True",
		Cfg.Database.User,
		Cfg.Database.Password,
		Cfg.Database.Name)
	log.Println("Trying to connect to dababase: ", connectionStr)
	var err error
	DB, err = gorm.Open("mysql", connectionStr)
	if err != nil {
		log.Panicf("Got error when connect database, the error is '%v'", err)
	}
	createTables()
}

func DBClose() {
	DB.Close()
}

// Init default database by dropping recreating tables with default data
func InitDefaultDatabaseData() {
	dropTables()
	createDefaultData()
}

func createTables() {
	DB.AutoMigrate(User{})
	DB.AutoMigrate(League{})
	DB.AutoMigrate(Game{})
	DB.AutoMigrate(GameResult{})
}

func dropTables() {
	log.Print("Dropping database tables...")
	defer log.Println("done")
	DB.DropTable(User{})
	DB.DropTable(League{})
	DB.DropTable(Game{})
	DB.DropTable(GameResult{})
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
	DB.Save(&user1)

	user2 := User{
		Login:     "robin",
		Password:  "robin",
		Name:      "robin penea",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	DB.Save(&user2)

	league := League{
		Name:      "league-test",
		Users:     []User{user1, user2},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	DB.Save(&league)
}
