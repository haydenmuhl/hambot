package chat

import (
	"log"

	"github.com/haydenmuhl/hambot/database"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/irc.v2"
)

func config(handler irc.Handler) irc.ClientConfig {
	db, err := database.Handle()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT username, password FROM bot_account WHERE id = 1")
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
