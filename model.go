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
	Cups   []Cup  `json:"-"`
	Games     []Game    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
}

type Cup struct {
	Id        int64  `json:"id"`
	Name      string `json:"name" sql:"size:255"`
	Type     string `json:"type" sql:"not null"`
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
	CupId int64  `json:"leagueId"`
	Cup   Cup `json:"-"`
	Users    []User
	Scores  []Score
}

type Score struct {
	Id             int64   `json:"id"`
	UserId         int64   `json:"userId"`
	Type     string `json:"type" sql:"not null"`
	CashGameScore float64 `json:"value,omitempty"`
	SitAndGoScore int32   `json:"value,omitempty"`
}

type UserCup struct {
	Id       int64
	UserId   int64
	User     User
	CupId int64
	Cup   Cup
}

type UserGame struct {
	Id     int64
	UserId int64
	User   User
	GameId int64
	Game   Game
}
