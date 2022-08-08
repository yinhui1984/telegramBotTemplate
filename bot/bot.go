package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func getToken() string {

	// 如果使用 heroku 运行程序 可以直接使用 os.Getenv("xxx") 获取
	// 如果直接运行程序 则可以使用这个库"github.com/joho/godotenv" 来解析.env文件然后使用 os.Getenv("xxx") 获取

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")
	if "" == token {
		log.Fatal("TOKEN is not set")
	}
	return token
}

var botApi *tgbotapi.BotAPI = nil

func GetBotInstance() *tgbotapi.BotAPI {

	if nil == botApi {
		b, err := tgbotapi.NewBotAPI(getToken())
		if err != nil {
			log.Panic("Error when create new bot:", err)
		}

		//true显示更多debug信息
		b.Debug = true

		log.Println("bot created,  username:", b.Self.UserName)
		botApi = b
	}

	return botApi
}
