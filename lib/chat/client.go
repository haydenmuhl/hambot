package chat

import (
	"log"
	"net"
	"strings"

	"gopkg.in/irc.v2"
)

func Client() *irc.Client {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		log.Fatalln(err)
	}
	var config irc.ClientConfig = config(newHandler())

	return irc.NewClient(conn, config)
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
