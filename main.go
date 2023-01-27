package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gogpt "github.com/sashabaranov/go-gpt3"
)

var conversationContext string

func main() {

	//telegram token
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))

	if err != nil {
		log.Panic(err)
		fmt.Println("MISSING_TELEGRAM_BOT_TOKEN")
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		//openai api
		var prompt string
		c := gogpt.NewClient(os.Getenv("OPENAI_API"))
		ctx := context.Background()

		if conversationContext == "" {
			prompt = update.Message.Text
		} else {
			prompt = conversationContext + update.Message.Text
		}
		req := gogpt.CompletionRequest{
			Model:            gogpt.GPT3TextDavinci003,
			MaxTokens:        150,
			Temperature:      0.9,
			TopP:             1,
			FrequencyPenalty: 0.0,
			PresencePenalty:  0.6,

			Prompt: prompt,
		}
		resp, err := c.CreateCompletion(ctx, req)
		if err != nil {
			return
		}
		if update.Message != nil { // jika mendapat pesan

			if update.Message.Text == "/start" {
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

			} else if update.Message.Text == "/help" {
				ctx := context.Background()
				req := gogpt.CompletionRequest{
					Model:            gogpt.GPT3TextDavinci003,
					MaxTokens:        50,
					Temperature:      0.9,
					TopP:             1,
					FrequencyPenalty: 0.0,
					PresencePenalty:  0.6,

					Prompt: "apa yang bisa chatGPT lakukan?",
				}
				resp, err := c.CreateCompletion(ctx, req)
				if err != nil {
					return
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, resp.Choices[0].Text)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			} else {

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, resp.Choices[0].Text)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}

		}
	}

}
