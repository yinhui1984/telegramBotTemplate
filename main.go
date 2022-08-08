package main

import (
	"BasicTelBot/bot"
	"BasicTelBot/commandsHandler"
	"BasicTelBot/events"

	//https://go-telegram-bot-api.dev
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {

	start := commandsHandler.StartCommand{}
	start.RegisterUpdaterEvent()

	mushenji := commandsHandler.MuShenJiCommand{}
	mushenji.RegisterUpdaterEvent()

	//----

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updatesChan := bot.GetBotInstance().GetUpdatesChan(updateConfig)

	log.Println("Enter Messages Loop")

	for update := range updatesChan {

		if update.Message != nil {
			if update.Message.IsCommand() {
				events.UpdaterEventInstance().Fire(update.Message)
			}
		} else if update.CallbackQuery != nil {
			events.UpdaterEventInstance().Fire(update.CallbackQuery)
		}

	}
}
