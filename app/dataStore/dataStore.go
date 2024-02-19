package dataStore

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var db *sql.DB

func GetDB() (*sql.DB, error) {
	cfg := getConfig()
	if db == nil {
		var err error
		db, err = sql.Open("sqlite3", cfg.DbFile)
		if err != nil {
			return nil, err
		}
		defer db.Close()
	}

	return db, nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
