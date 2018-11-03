package main

import (
	"log"

	"github.com/haydenmuhl/hambot/admin"
	"github.com/haydenmuhl/hambot/database"
)

func main() {
	var err error

	err = database.Init()
	if err != nil {
		log.Println("Unable to initialize the database.")
		log.Fatalln(err)
	}

	err = admin.Server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
