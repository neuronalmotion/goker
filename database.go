package goker

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // used by gorm
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func init() {
	connectionStr := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True",
		GokerCtx.Cfg.Database.User,
		GokerCtx.Cfg.Database.Password,
		GokerCtx.Cfg.Database.Name)
	log.Println("Trying to connect to dababase: ", connectionStr)
	var err error
	GokerCtx.DB, err = gorm.Open("mysql", connectionStr)
	if err != nil {
		log.Panicf("Got error when connect database, the error is '%v'", err)
	}
	createTables()
}

func DBClose() {
	GokerCtx.DB.Close()
}

func DBClear() {
	log.Print("Dropping database tables...")
	defer log.Println("done")
	GokerCtx.DB.DropTable(User{})
	GokerCtx.DB.DropTable(Cup{})
	GokerCtx.DB.DropTable(Game{})
	GokerCtx.DB.DropTable(Score{})
	GokerCtx.DB.DropTable(UserCup{})
	GokerCtx.DB.DropTable(UserGame{})
}

// Init default database by dropping recreating tables with default data
func DBDefaultData() {
	DBClear()
	createDefaultData()
}

func FillCupData(cup *Cup) {
	GokerCtx.DB.Model(cup).Related(&cup.Owner, "OwnerId")
	cup.Users = DBGetUsersForCup(cup.Id)
	GokerCtx.DB.Model(cup).Related(&cup.Games)
}

func DBGetCupsForUser(userId int64) []Cup {
	cups := []Cup{}
	GokerCtx.DB.Debug().Raw("SELECT l.id, l.name, l.owner_id, l.created_at, l.updated_at FROM cups l, user_cups WHERE l.id = user_cups.cup_id AND user_cups.user_id = ?", userId).Scan(&cups)
	for i := 0; i < len(cups); i++ {
		FillCupData(&cups[i])
	}
	return cups
}

func DBGetUsersForCup(cupId int64) []User {
	users := []User{}
	GokerCtx.DB.Raw("SELECT u.id, u.login, u.password, u.email, u.name, u.created_at, u.updated_at, u.deleted_at FROM users u, user_cups ul WHERE u.id = ul.user_id AND ul.cup_id = ?", cupId).Scan(&users)
	return users
}

func DBGetUsersForGame(gameId int64) []User {
	users := []User{}
	GokerCtx.DB.Raw("SELECT u.id, u.login, u.password, u.email, u.name, u.created_at, u.updated_at, u.deleted_at FROM users u, user_games ug WHERE u.id = ug.user_id AND ug.game_id = ?", gameId).Scan(&users)
	return users
}

func DBGetGamesForUser(userId int64) []Game {
	games := []Game{}
	GokerCtx.DB.Debug().Raw("SELECT g.id, g.type, g.cup_id FROM games g, user_games ug WHERE g.id = ug.game_id AND ug.user_id = ?", userId).Scan(&games)
	return games
}

func createTables() {
	GokerCtx.DB.AutoMigrate(User{})
	GokerCtx.DB.AutoMigrate(Cup{})
	GokerCtx.DB.AutoMigrate(Game{})
	GokerCtx.DB.AutoMigrate(Score{})
	GokerCtx.DB.AutoMigrate(UserCup{})
	GokerCtx.DB.AutoMigrate(UserGame{})
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
	GokerCtx.DB.Save(&user1)

	user2 := User{
		Login:     "robin",
		Password:  "robin",
		Name:      "robin penea",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	GokerCtx.DB.Save(&user2)

	cup := Cup{
		Name:      "cup-test",
		Owner:     user1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	GokerCtx.DB.Save(&cup)

	GokerCtx.DB.Save(&UserCup{User: user1, Cup: cup})
	GokerCtx.DB.Save(&UserCup{User: user2, Cup: cup})

	// Game data
	game := Game{Type: GameTypeCashGame, Cup: cup}
	GokerCtx.DB.Save(&game)

	GokerCtx.DB.Save(&UserGame{User: user1, Game: game})
	GokerCtx.DB.Save(&UserGame{User: user2, Game: game})
}
