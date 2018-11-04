package chat

import (
	"log"

	"github.com/haydenmuhl/hambot/database"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/irc.v2"
)

func config(handler irc.Handler) irc.ClientConfig {
	db := database.Handle()

	rows, err := db.Query("SELECT username, password FROM bot_credentials WHERE id = 1")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	rows.Next()
	var username string
	var password string
	err = rows.Scan(&username, &password)
	if err != nil {
		log.Fatalln(err)
	}

	return irc.ClientConfig{
		Nick:    username,
		Pass:    password,
		User:    username,
		Handler: handler,
	}
}

func HasCredentials() bool {
	db := database.Handle()

	rows, err := db.Query("SELECT count(id) FROM bot_credentials")
	defer rows.Close()
	if err != nil {
		return false
	}
	if !rows.Next() {
		return false
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}
