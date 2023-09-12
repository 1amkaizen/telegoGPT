package controllers

import (
	
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/1amkaizen/telegoGPT/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func AccessOpenAIAPI(prompt string) (string, error) {
	c := gogpt.NewClient(os.Getenv("OPENAI_API"))
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        250,
		Temperature:      0.9,
		TopP:             1,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.6,

		Prompt: prompt,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Text, nil
}

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
	response, err := AccessOpenAIAPI(update.Message.Text)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("[%s] %s %s", update.Message.From.UserName, update.Message.Text, response)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)

}

func HandleStartCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("UserName :%s", update.Message.From.UserName)
	log.Printf("ID :%d", update.Message.Chat.ID)
	log.Printf("Text: %s", update.Message.Text)

	// button
	twitterButton := tgbotapi.NewInlineKeyboardButtonData("Twitter🐦", "twitter")
	githubButton := tgbotapi.NewInlineKeyboardButtonData("Github🐙", "github")
	railwayButton := tgbotapi.NewInlineKeyboardButtonData("Railway🚂", "railway")
	replitButton := tgbotapi.NewInlineKeyboardButtonData("Replit🚀", "replit")
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			twitterButton,
			githubButton,
		),
		tgbotapi.NewInlineKeyboardRow(
			railwayButton,
			replitButton,
		),
	)

	
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hallo, @"+update.Message.From.UserName+"! Selamat datang di bot saya, bagaimana saya bisa membantumu hari ini?")
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = keyboard
	bot.Send(msg)

	// send message to me
	SECRET := os.Getenv("SECRET")
	SECRET64, _ := strconv.ParseInt(SECRET, 10, 64)
	msgToYou := tgbotapi.NewMessage(SECRET64, "User @"+update.Message.From.UserName+" with ID:"+strconv.FormatInt(update.Message.Chat.ID, 10)+" masuk")

	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msgToYou)

	// Add user to database
	user := &models.Users{
		UserID:   strconv.FormatInt(update.Message.Chat.ID, 10),
		UserName: update.Message.From.UserName,
	}

	var existingUser models.Users
	if err := models.DB.Where("user_id = ?", user.UserID).First(&existingUser).Error; err == nil {
		fmt.Println("User sudah terdaftar")
		return
	}

	if err := models.DB.Create(user).Error; err != nil {
		fmt.Println("Gagal menyimpan data user")
		return
	}

	models.DB.Create(user)

	var userFromDB models.Users
	models.DB.First(&userFromDB, "user_id = ?", user.UserID)
	log.Println("User from database:", userFromDB)

}

func HandleHelpCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	c := gogpt.NewClient(os.Getenv("OPENAI_API"))

	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        150,
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


