package main

import (
        "context"
        "log"
        "os"

        tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
        gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {

        //telegram token
        bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))

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

                c := gogpt.NewClient(os.Getenv("OPENAI_API"))
                ctx := context.Background()
                req := gogpt.CompletionRequest{
                        Model: gogpt.GPT3TextDavinci003,
                        MaxTokens: 999,

                        Prompt: update.Message.Text,
                }
                resp, err := c.CreateCompletion(ctx, req)
                if err != nil {
                        return
                }
                if update.Message.Text == "/start" {
                        log.Printf("UserName :%s", update.Message.From.UserName)
                        log.Printf("ID :%d", update.Message.Chat.ID)
                        log.Printf("Text: %s", update.Message.Text)
                        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hallo,Selamat datang di telegoGPT")
                        bot.Send(msg)
                        
                        // send message to you
                        msgToYou := tgbotapi.NewMessage(2116777065, "User "+update.Message.From.UserName+" with ID:"+strconv.FormatInt(update.Message.Chat.ID,10)+" just use /start command")
                        bot.Send(msgToYou)

                } else if update.Message.Text == "/help" {
                        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Tanyakan apa saja atau beri perintah apa saja.\n\nContoh penggunaan : \n\n-Siapa presiden indonesia pertama?\n-Buat deskripsi makanan ringan.")
                        bot.Send(msg)

                } else if update.Message != nil { // jika mendapat pesan
                        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

                        //get user chat id
                       chat, err := bot.GetChat(tgbotapi.ChatConfigWithUser{UserID: update.Message.Chat.ID})
                        if err != nil {
                                log.Println(err)
                                continue
                        }
                        //get user id
                        var userId int64
                        userId = chat.ID
                        //send message to user
                        msg := tgbotapi.NewMessage(userId, resp.Choices[0].Text)
                        bot.Send(msg)
                }
        }
}
