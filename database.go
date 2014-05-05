package gocker

import (
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB gorm.DB

func init() {
    username := "root"
    password := ""
    database := "neuronalpoker"
    connectionStr := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True", username, password, database)
    fmt.Println("Trying to connect to dababase: ", connectionStr)
    var err error
    DB, err = gorm.Open("mysql", connectionStr)
    if err != nil {
        panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
    }
	createTables()
}

func DBClose() {
    DB.Close()
}

func createTables() {
	DB.AutoMigrate(User{})
	DB.AutoMigrate(League{})
	DB.AutoMigrate(Game{})
	DB.AutoMigrate(GameResult{})
}

func dropTables() {
    fmt.Println("Dropping database tables...")
	DB.DropTable(User{})
	DB.DropTable(League{})
	DB.DropTable(Game{})
	DB.DropTable(GameResult{})
}

func createDefaultData() {
    fmt.Println("Creating default database data...")
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
func InitDefaultDatabaseData() {
	dropTables()
	createDefaultData()
}
