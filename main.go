package main

import (
	"fmt"
	"log"
	"os"

	"telegoGPT/controllers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var conversationContext string

func main() {

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
		var prompt string

		if conversationContext == "" {
			prompt = update.Message.Text
		} else {
			prompt = conversationContext + update.Message.Text
		}
		resp, err := controllers.AccessOpenAIAPI(prompt)
		if err != nil {
			log.Println("error: ", err)
		}
		return resp.Choices[0].Text, nil

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
