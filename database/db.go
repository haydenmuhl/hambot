package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var handle *sql.DB

func Handle() (*sql.DB, error) {

	if handle == nil {
		var err error
		handle, err = connect()

		if err != nil {
			return nil, err
		}
	}
	return handle, nil
}

var (
	dbPath = fmt.Sprintf("%s/.hambot/hambot.db", os.Getenv("HOME"))
)

func connect() (*sql.DB, error) {
	return sql.Open("sqlite3", dbPath)
}
