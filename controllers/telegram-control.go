package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gogpt "github.com/sashabaranov/go-gpt3"
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
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        60,
		Temperature:      0.5,
		TopP:             0.3,
		FrequencyPenalty: 0.5,
		PresencePenalty:  0.0,

		Prompt: "apa yang bisa chatGPT lakukan?",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, resp.Choices[0].Text)
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)

}
