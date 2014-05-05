package gocker

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

func createTables(db *gorm.DB) {
    db.AutoMigrate(User{})
    db.AutoMigrate(League{})
    db.AutoMigrate(Game{})
    db.AutoMigrate(GameResult{})
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
}

