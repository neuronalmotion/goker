package gocker

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func createTables(db *gorm.DB) {
	db.AutoMigrate(User{})
	db.AutoMigrate(League{})
	db.AutoMigrate(Game{})
	db.AutoMigrate(GameResult{})
}

func createDefaultData(db *gorm.DB) {
    db.Debug().Unscoped().Where("login = ?", "guillaume").Delete(User{})
    db.Debug().Unscoped().Where("login = ?", "robin").Delete(User{})
    db.Debug().Unscoped().Where("name = ?", "league-test").Delete(League{})

	user1 := User{
		Login:     "guillaume",
		Password:  "guillaume",
		Name:      "guillaume lazar",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	db.Save(&user1)

    user2 := User{
		Login:     "robin",
		Password:  "robin",
		Name:      "robin penea",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	db.Save(&user2)

    league := League{
        Name: "league-test",
        Users: []User{user1, user2},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	db.Save(&league)


}
func DatabaseInit() {
	username := "root"
	password := ""
	database := "neuronalpoker"
	connectionStr := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True", username, password, database)
	fmt.Println("Trying to connect to dababase: ", connectionStr)
	db, err := gorm.Open("mysql", connectionStr)

	fmt.Println("Database connection result: ", err)
	d := db.DB()
	fmt.Println("Ping result:", d.Ping())

	createTables(&db)
    createDefaultData(&db)
}
