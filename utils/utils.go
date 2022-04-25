package utils

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/infiniteloopcloud/discord-jira/env"
	"github.com/infiniteloopcloud/discord-jira/jira"
)

var session *discordgo.Session
var channelsCache map[string]string

func GetEvent(raw []byte) (jira.IssueWrapper, error) {
	var issue jira.IssueWrapper
	err := json.Unmarshal(raw, &issue)
	if err != nil {
		return jira.IssueWrapper{}, err
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
		channels, err := GetSession().GuildChannels(env.Configuration().BotGuild)
		if err != nil {
			log.Print(err)
		}
		for _, channel := range channels {
			if name == NormalizeChannelName(channel.Name) {
				channelsCache[NormalizeChannelName(channel.Name)] = channel.ID
				return channel.ID
			}
		}
	}
	return ""
}

func GetSession() *discordgo.Session {
	if session == nil {
		var err error
		session, err = discordgo.New("Bot " + env.Configuration().BotToken)
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
		}
	}
	return session
}

func NormalizeChannelName(str string) string {
	return strings.ToLower(strings.Replace(str, " ", "-", -1))
}
