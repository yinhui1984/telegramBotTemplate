package utils

import (
	"BasicTelBot/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func SendMessage(chatId int64, text string, isHtml bool, replayMarkup interface{}) {
	reply := tgbotapi.NewMessage(chatId, text)

	if isHtml {
		reply.ParseMode = "html"
	}

	if nil != replayMarkup {
		reply.ReplyMarkup = replayMarkup
	}

	_, err := bot.GetBotInstance().Send(reply)
	if err != nil {
		log.Println("error when send message:", err)
	}
}

// RequestCallback 相当于ACK
func RequestCallback(query *tgbotapi.CallbackQuery) {
	callback := tgbotapi.NewCallback(query.ID, query.Data)
	_, err := bot.GetBotInstance().Request(callback)
	if err != nil {
		return
	}
}
