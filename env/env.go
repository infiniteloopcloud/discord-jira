package env

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var (
	configuration *Static
	PATH          string = "JIRA_BOT_CONFIG"
)

type Static struct {
	SkipRepoPushMessages string `json:"skip_repo_push_message"`
	BotToken             string `json:"bot_token"`
	BotGuild             string `json:"bot_guild"`
	Address              string `json:"address"`
}

func Configuration() *Static {
	path := os.Getenv(PATH)
	if configuration == nil {
		// read from path
		var s Static
		// json unmarshal into s
		configuration = &s

		file, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
		}

		json.Unmarshal(file, &configuration)

	}
	return configuration
}
