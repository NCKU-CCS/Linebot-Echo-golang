package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Configuration struct {
	CHANNEL_SECRET       string
	CHANNEL_ACCESS_TOKEN string
}

func loadConfigFile(fPath string) *Configuration {
	configuration := Configuration{}

	file, _ := os.Open(fPath)
	defer file.Close()

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return &configuration
}

func main() {

	// load configuration
	botConfigs := loadConfigFile("./bot_settings.json")
	client := &http.Client{}
	bot, err := linebot.New(botConfigs.CHANNEL_SECRET, botConfigs.CHANNEL_ACCESS_TOKEN, linebot.WithHTTPClient(client))
	fmt.Println(client)
}
