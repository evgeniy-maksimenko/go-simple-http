package models

import (
	"database/sql"
	_ "github.com/gotk/pg"
	"log"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	db,err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func NewDB(dataSourceName string) (*sql.DB, error) {
	db,err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db, nil
}

