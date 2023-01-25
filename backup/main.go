package main

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	//telegram token
	bot, err := tgbotapi.NewBotAPI("5928666195:AAFo5iRKICjPbE74xNa2DPBoOnFd9Xv5Ybo")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		//openai api
		c := gogpt.NewClient("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")
		ctx := context.Background()
		req := gogpt.CompletionRequest{
			Model:     gogpt.GPT3TextDavinci001,
			MaxTokens: 999,
			Prompt:    update.Message.Text,
		}
		resp, err := c.CreateCompletion(ctx, req)
		if err != nil {
			return
		}

		if update.Message != nil { // jika mendapat pesan
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, resp.Choices[0].Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
		//
	}
}
