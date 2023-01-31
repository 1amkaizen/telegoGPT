package main

import (
	"fmt"
	"log"
	"os"

	"github.com/1amkaizen/telegoGPT/controllers"
	"github.com/1amkaizen/telegoGPT/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	models.ConnectDatabase()

	//telegram token
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))

	if err != nil {
		log.Panic(err)
		fmt.Println("MISSING_TELEGRAM_BOT_TOKEN")
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {

		//openai api
		if update.Message != nil { // jika mendapat pesan

			if update.Message.Text == "/start" {

				controllers.HandleStartCommand(bot, update)

			} else if update.Message.Text == "/help" {
				controllers.HandleHelpCommand(bot, update)

			} else {
				controllers.SendMessage(bot, update)

			}

		}
	}
}
