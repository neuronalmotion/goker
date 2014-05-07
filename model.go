package gocker

import (
	_ "database/sql"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	Login     string    `json:"login" sql:"size:255;unique"`
	Password  string    `json:"-" sql:"size:255"`
	Email     string    `json:"email" sql:"type:varchar(100)"`
	Name      string    `json:"name" sql:"size:255"`
	Leagues   []League  `json:"-"`
	Games     []Game    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
}

type League struct {
	Id        int64  `json:"id"`
	Name      string `json:"name" sql:"size:255"`
	OwnerId   int64  `json:"-"`
	Owner     User
	Users     []User
	Games     []Game
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

const GameTypeSitAndGo = "sitandgo"
const GameTypeCashGame = "cashgame"

type Game struct {
	Id       int64  `json:"id"`
	Type     string `json:"type" sql:"not null"`
	LeagueId int64  `json:"leagueId"`
	League   League `json:"-"`
	Users    []User
	Results  []GameResult
}

type GameResult struct {
	Id             int64   `json:"id"`
	UserId         int64   `json:"userId"`
	CashGameResult float64 `json:"result,omitempty"`
	SitAndGoResult int32   `json:"result,omitempty"`
}

type UserLeague struct {
	Id       int64
	UserId   int64
	User     User
	LeagueId int64
	League   League
}

type UserGame struct {
	Id     int64
	UserId int64
	User   User
	GameId int64
	Game   Game
}
