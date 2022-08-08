package commandsHandler

import (
	"BasicTelBot/events"
	"BasicTelBot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	MushenjiCommandName = "mushenji"
)

var keyboardMuShenJi = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("0001_天黑别出门", "0001_天黑别出门.mp3"),
		tgbotapi.NewInlineKeyboardButtonData("0002_四灵血", "0002_四灵血.mp3"),
		tgbotapi.NewInlineKeyboardButtonData("0003_神通", "0003_神通.mp3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("0004_天魔造化功", "0004_天魔造化功.mp3"),
		tgbotapi.NewInlineKeyboardButtonData("0005_漓江五老", "0005_漓江五老.mp3"),
	),
)

type MuShenJiCommand struct {
}

func (c MuShenJiCommand) showChapters(query *tgbotapi.CallbackQuery) {
	utils.RequestCallback(query)
	utils.SendMessage(query.Message.Chat.ID, "<b>请选择章节</b>", true, keyboardMuShenJi)
}

func (c MuShenJiCommand) showChapters2(chatId int64) {
	utils.SendMessage(chatId, "<b>请选择章节</b>", true, keyboardMuShenJi)
}

func (c MuShenJiCommand) play(query *tgbotapi.CallbackQuery) {
	utils.RequestCallback(query)
	utils.SendMessage(query.Message.Chat.ID, query.Data+" playing...", false, nil)
}

func (c MuShenJiCommand) RegisterUpdaterEvent() {
	events.UpdaterEventInstance().Register(func(event *events.UpdaterEvent, arg interface{}) {

		// if the type of msg is *tgbotapi.Message
		if msg, ok := arg.(*tgbotapi.Message); ok {
			switch msg.Command() {
			case MushenjiCommandName: // "/mushenji"
				c.showChapters2(msg.Chat.ID)
			default:
				break
			}
		}

		// if the type of msg is *tgbotapi.CallbackQuery
		if query, ok := arg.(*tgbotapi.CallbackQuery); ok {
			switch query.Data {
			case MushenjiCommandName:
				c.showChapters(query)
			case "0001_天黑别出门.mp3":
				c.play(query)
			case "0002_四灵血.mp3":
				c.play(query)
			case "0004_天魔造化功.mp3":
				c.play(query)
			case "0005_漓江五老.mp3":
				c.play(query)
			}
		}
	})
}
