package main

import (
	"log"
	"os"

	runner "github.com/infiniteloopcloud/jira-dc-bot/runner"
)

const (
	LOGFILE = "jira-dc-bot.log"
)

func init() {
	// Set the log file
	file, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func main() {
	runner.Run()
}
