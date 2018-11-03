package database

import (
	"database/sql"
)

type migration struct {
	Version string
	Query   string
}

func (m migration) ApplyTo(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	row := tx.QueryRow("SELECT * FROM migrations WHERE version = ?", m.Version)
	var version string
	err = row.Scan(&version)
	if err == nil {
		return nil
	}

	_, err = tx.Exec(m.Query)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO migrations (version) VALUES (?)", m.Version)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func migrate(db *sql.DB) error {
	var err error
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS migrations (version TEXT UNIQUE)`)
	if err != nil {
		return err
	}

	for _, migration := range migrations {
		err = migration.ApplyTo(db)
		if err != nil {
			return err
		}
	}
	return nil
}

var migrations = []migration{
	{
		Version: "2018110200",
		Query: `CREATE TABLE bot_credentials (
				  id INTEGER PRIMARY KEY AUTOINCREMENT,
				  username TEXT,
				  password TEXT
				);`,
	},
}
