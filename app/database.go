package app

import (
	"database/sql"
	"time"

	"github.com/dihanto/go-mastermind/helper"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go_mastermind")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(10)

	return db

}
