package main

import (
	"log"

	"github.com/haydenmuhl/hambot/chat"
)

func main() {
	err := chat.Client().Run()

	if err != nil {
		log.Fatalln(err)
	}
}
