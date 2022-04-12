package runner

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/infiniteloopcloud/jira-dc-bot/env"
	handler "github.com/infiniteloopcloud/jira-dc-bot/handler"
)

var session *discordgo.Session
var channelsCache map[string]string

func getChannelID(name string) string {
	if channelsCache == nil {
		channelsCache = make(map[string]string)
	}
	if id, ok := channelsCache[name]; ok {
		return id
	} else {
		channels, err := getSession().GuildChannels(os.Getenv(env.GuildID))
		if err != nil {
			log.Print(err)
		}
		for _, channel := range channels {
			if name == channel.Name {
				channelsCache[channel.Name] = channel.ID
				return channel.ID
			}
		}
	}
	return ""
}

func getSession() *discordgo.Session {
	if session == nil {
		var err error
		session, err = discordgo.New("Bot " + os.Getenv(env.Token))
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
		}
	}
	return session
}

func hello(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}

	event, err := handler.GetEvent(body)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}

	channel, message, err := handler.Handle(event.EventType, body)
	if err != nil {
		log.Printf("[ERROR] [%s] %s", event, err.Error())
		return
	}
	channelID := getChannelID(channel)
	if channelID == "" {
		channelID = getChannelID("unknown")
	}
	if channelID != "" && message != nil {
		_, err = getSession().ChannelMessageSendEmbed(channelID, message)
		if err != nil {
			log.Printf("[ERROR] [%s] %s", event, err.Error())
		}
	}

	fmt.Fprintf(w, "ACK")
}

func Run() {
	address := ":8080"
	if a := os.Getenv(env.Address); a != "" {
		address = a
	}

	http.HandleFunc("/webhook", hello)
	log.Printf("Server listening on %s", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}
}
