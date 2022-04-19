package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/infiniteloopcloud/discord-jira/env"
	handler "github.com/infiniteloopcloud/discord-jira/handler"
)


var session *discordgo.Session
var channelsCache map[string]string

func GetEvent(raw []byte) (handler.Issue, error) {
	var issue handler.Issue
	err := json.Unmarshal(raw, &issue)
	if err != nil {
		return handler.Issue{}, err
	}
	return issue, nil
}

func GetChannelID(name string) string {
	if channelsCache == nil {
		channelsCache = make(map[string]string)
	}
	if id, ok := channelsCache[name]; ok {
		return id
	} else {
		channels, err := GetSession().GuildChannels(os.Getenv(env.GuildID))
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

func GetSession() *discordgo.Session {
	if session == nil {
		var err error
		session, err = discordgo.New("Bot " + os.Getenv(env.Token))
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
		}
	}
	return session
}
