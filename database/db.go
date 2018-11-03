package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var handle *sql.DB

func Handle() *sql.DB {
	return handle
}

var (
	dbPath = fmt.Sprintf("%s/.hambot/hambot.db", os.Getenv("HOME"))
)

func connect() (*sql.DB, error) {
	return sql.Open("sqlite3", dbPath)
}

// Initialize the database. If the database doesn't exist, create the database
func Init() error {
	var err error
	if handle == nil {
		handle, err = connect()

		if err != nil {
			return err
		}
	}

	err = migrate(handle)
	if err != nil {
		return err
	}

	return nil
}
