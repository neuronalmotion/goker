package gocker

import (
    _ "database/sql"
    "time"
)

type User struct {
    Id int64
    Login string `sql:"size:255"`
    Password string `sql:"size:255"`
    Email string `sql:"type:varchar(100)"`
    Name string `sql:"size:255"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

type League struct {
    Id int64
    Name string `sql:"size:255"`
    Users []User
    Games []Game
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Game struct {
    Id int64
    Type string `sql:"not null"`
    Users []User
    Results []GameResult
}

type GameResult struct {
    Id int64
    UserId int64
    CashGameResult float64
    SitAndGoResult int32
}

