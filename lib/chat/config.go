package chat

import (
	"log"
	"strings"

	"github.com/haydenmuhl/hambot/lib/database"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/irc.v2"
)

type botConfig struct {
	username, password, channel string
}

func (b *botConfig) IrcConfig() irc.ClientConfig {
	return irc.ClientConfig{
		Nick:    b.username,
		Pass:    b.password,
		User:    b.username,
		Handler: &Handler{strings.ToLower(b.channel)},
	}
}

func config() *botConfig {
	db := database.Handle()

	query := `SELECT cr.username, cr.password, ch.name
			  FROM channel ch
			  INNER JOIN credential cr on ch.credential_id = cr.id
			  WHERE ch.id = 1`

	row := db.QueryRow(query)

	bot := botConfig{}
	err := row.Scan(&bot.username, &bot.password, &bot.channel)
	if err != nil {
		log.Fatalln(err)
	}
	return &bot
}

func HasCredentials() bool {
	db := database.Handle()

	rows, err := db.Query("SELECT count(id) FROM credential")
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
