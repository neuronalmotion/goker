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
	GockerCtx.DB.DropTable(UserLeague{})
	GockerCtx.DB.DropTable(UserGame{})
}

// Init default database by dropping recreating tables with default data
func DBDefaultData() {
	DBClear()
	createDefaultData()
}

func FillLeagueData(league *League) {
	GockerCtx.DB.Model(league).Related(&league.Owner, "OwnerId")
	league.Users = DBGetUsersForLeague(league.Id)
	GockerCtx.DB.Model(league).Related(&league.Games)
}

func DBGetLeaguesForUser(userId int64) []League {
	leagues := []League{}
	GockerCtx.DB.Debug().Raw("SELECT l.id, l.name, l.owner_id, l.created_at, l.updated_at FROM leagues l, user_leagues WHERE l.id = user_leagues.league_id AND user_leagues.user_id = ?", userId).Scan(&leagues)
	for i := 0; i < len(leagues); i++ {
		FillLeagueData(&leagues[i])
	}
	return leagues
}

func DBGetUsersForLeague(leagueId int64) []User {
	users := []User{}
	GockerCtx.DB.Raw("SELECT u.id, u.login, u.password, u.email, u.name, u.created_at, u.updated_at, u.deleted_at FROM users u, user_leagues ul WHERE u.id = ul.user_id AND ul.league_id = ?", leagueId).Scan(&users)
	return users
}

func DBGetUsersForGame(gameId int64) []User {
	users := []User{}
	GockerCtx.DB.Raw("SELECT u.id, u.login, u.password, u.email, u.name, u.created_at, u.updated_at, u.deleted_at FROM users u, user_games ug WHERE u.id = ug.user_id AND ug.game_id = ?", gameId).Scan(&users)
	return users
}

func DBGetGamesForUser(userId int64) []Game {
	games := []Game{}
	GockerCtx.DB.Debug().Raw("SELECT g.id, g.type, g.league_id FROM games g, user_games ug WHERE g.id = ug.game_id AND ug.user_id = ?", userId).Scan(&games)
	return games
}

func createTables() {
	GockerCtx.DB.AutoMigrate(User{})
	GockerCtx.DB.AutoMigrate(League{})
	GockerCtx.DB.AutoMigrate(Game{})
	GockerCtx.DB.AutoMigrate(GameResult{})
	GockerCtx.DB.AutoMigrate(UserLeague{})
	GockerCtx.DB.AutoMigrate(UserGame{})
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
		Owner:     user1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	GockerCtx.DB.Save(&league)

	GockerCtx.DB.Save(&UserLeague{User: user1, League: league})
	GockerCtx.DB.Save(&UserLeague{User: user2, League: league})

	// Game data
	game := Game{Type: GameTypeCashGame, League: league}
	GockerCtx.DB.Save(&game)

	GockerCtx.DB.Save(&UserGame{User: user1, Game: game})
	GockerCtx.DB.Save(&UserGame{User: user2, Game: game})
}
