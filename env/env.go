package env

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var (
	configuration *Static
	ConfigFlag    string = "JIRA_BOT_CONFIG"
)

type Static struct {
	BaseURL  string `json:"base_url"`
	BotToken string `json:"bot_token"`
	BotGuild string `json:"bot_guild"`
	Address  string `json:"address"`
}

func Configuration() *Static {
	if configuration == nil {
		var path string
		if path = os.Getenv(ConfigFlag); path == "" {
			path = "./config.json"
		}

		file, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("[ERROR] %s\n", err.Error())
			return nil
		}
		var s Static
		if err := json.Unmarshal(file, &s); err != nil {
			log.Printf("[ERROR] unmarshal file: %s", err.Error())
			return nil
		}
		configuration = &s
	}
	return configuration
}

func (s Static) Dump() {
	log.Printf("[INFO] Configuration:\n")
	log.Printf("[INFO] \tBaseURL: %s\n", s.BaseURL)
	log.Printf("[INFO] \tBotToken: %s\n", s.BotToken)
	log.Printf("[INFO] \tBotGuild: %s\n", s.BotGuild)
	log.Printf("[INFO] \tAddress: %s\n", s.Address)
}
