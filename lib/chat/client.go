package chat

import (
	"fmt"
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
	bot := config()

	return irc.NewClient(conn, bot.IrcConfig())
}

type Handler struct {
	channel string
}

func (h *Handler) Handle(client *irc.Client, message *irc.Message) {
	log.Println(message)
	if message.Command == "001" {
		client.Write(fmt.Sprintf("JOIN #%s", h.channel))
	} else if message.Command == "PING" {
		handlePing(client, message)
	} else if message.Command == "PRIVMSG" && strings.ToLower(message.Params[1]) == "bleep" {
		client.Write(fmt.Sprintf(":roboraccoon!roboraccoon@roboraccoon.tmi.twitch.tv PRIVMSG #%s :bloop", h.channel))
	}
}

func newHandler(channel string) *Handler {
	return &Handler{strings.ToLower(channel)}
}

func handlePing(c *irc.Client, m *irc.Message) {
	reply := m.Copy()
	reply.Command = "PONG"
	c.WriteMessage(reply)
}
