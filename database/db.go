package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var handle *sql.DB

func Handle() *sql.DB {
	return handle
}

var (
	dbDir  = fmt.Sprintf("%s/.hambot", os.Getenv("HOME"))
	dbFile = "hambot.db"
	dbPath = fmt.Sprintf("%s/%s", dbDir, dbFile)
)

func connect() (*sql.DB, error) {
	log.Printf("Opening database at %s\n", dbPath)
	return sql.Open("sqlite3", dbPath)
}

// Initialize the database. If the database doesn't exist, create the database
func Init() error {
	var err error
	err = os.MkdirAll(dbDir, os.ModeDir | 0755)

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
