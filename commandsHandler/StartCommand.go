package commandsHandler

import (
	"BasicTelBot/events"
	"BasicTelBot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	StartCommandName = "start"
	HelpQueryName    = "start_help"
)

type StartCommand struct {
}

func (c StartCommand) welcome(chatId int64) {
	utils.SendMessage(chatId, "<b>欢迎, 请选择功能</b>", true, tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("me", "https://github.com/yinhui1984"),
			tgbotapi.NewInlineKeyboardButtonData("帮助", HelpQueryName),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("牧神记", MushenjiCommandName),
		),
	))
}

func (c StartCommand) showHelp(query *tgbotapi.CallbackQuery) {
	utils.RequestCallback(query)
	utils.SendMessage(query.Message.Chat.ID, "这是 <ins>帮助</ins>文档, 支持的命令:\n /start\n /mushenji\n", true, nil)
}

func (c StartCommand) RegisterUpdaterEvent() {
	events.UpdaterEventInstance().Register(func(event *events.UpdaterEvent, arg interface{}) {

		// if the type of msg is *tgbotapi.Message
		if msg, ok := arg.(*tgbotapi.Message); ok {
			switch msg.Command() {
			case StartCommandName: // "/start"
				c.welcome(msg.Chat.ID)
			}
		}

		// if the type of msg is *tgbotapi.CallbackQuery
		if query, ok := arg.(*tgbotapi.CallbackQuery); ok {
			switch query.Data {
			case HelpQueryName:
				c.showHelp(query)
			case StartCommandName:
				c.welcome(query.Message.Chat.ID)
			}
		}
	})
}
