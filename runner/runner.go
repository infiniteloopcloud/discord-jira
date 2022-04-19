package runner

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/infiniteloopcloud/discord-jira/env"
	handler "github.com/infiniteloopcloud/discord-jira/handler"
	utils "github.com/infiniteloopcloud/discord-jira/utils"
)

func webhookHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}

	event, err := utils.GetEvent(body)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}

	channel, message, err := handler.Handle(event.EventType, body)
	if err != nil {
		log.Printf("[ERROR] [%s] %s", event, err.Error())
		return
	}
	channelID := utils.GetChannelID(channel)
	if channelID == "" {
		channelID = utils.GetChannelID("unknown")
	}
	fmt.Println(channel, channelID)
	if channelID != "" && message != nil {
		_, err = utils.GetSession().ChannelMessageSendEmbed(channelID, message)
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

	http.HandleFunc("/webhook", webhookHandler)
	log.Printf("Server listening on %s", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}
}
