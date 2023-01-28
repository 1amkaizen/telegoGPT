package controllers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetupBot() (*tgbotapi.BotAPI, error) {
	//telegram token
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))

	if err != nil {
		log.Panic(err)
		fmt.Println("MISSING_TELEGRAM_BOT_TOKEN")
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot, nil
}

func SendMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}

func HandleStartCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("UserName :%s", update.Message.From.UserName)
	log.Printf("ID :%d", update.Message.Chat.ID)
	log.Printf("Text: %s", update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hallo, @"+update.Message.From.UserName+"! Selamat datang di bot saya, bagaimana saya bisa membantumu hari ini?")
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)

	// send message to me
	msgToYou := tgbotapi.NewMessage(2116777065, "User @"+update.Message.From.UserName+" with ID:"+strconv.FormatInt(update.Message.Chat.ID, 10)+" masuk")
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msgToYou)

}
