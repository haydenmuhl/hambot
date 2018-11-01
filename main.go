package main

import (
	"log"
	"net"
	"strings"

	"gopkg.in/irc.v2"
)

func main() {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		log.Fatalln(err)
	}

	// config() returns an irc.ClientConfig with credentials
	// The credentials are hard coded into the function, and it is not checked in to the repo
	var config irc.ClientConfig = config(newHandler())

	client := irc.NewClient(conn, config)
	err = client.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func newHandler() irc.Handler {
	return irc.HandlerFunc(func(client *irc.Client, message *irc.Message) {
		log.Println(message)
		if message.Command == "001" {
			client.Write("JOIN #soonraccoon")
		} else if message.Command == "PING" {
			handlePing(client, message)
		} else if message.Command == "PRIVMSG" && strings.ToLower(message.Params[1]) == "bleep" {
			client.Write(":roboraccoon!roboraccoon@roboraccoon.tmi.twitch.tv PRIVMSG #soonraccoon :bloop")
		}
	})
}

func handlePing(c *irc.Client, m *irc.Message) {
	reply := m.Copy()
	reply.Command = "PONG"
	c.WriteMessage(reply)
}
