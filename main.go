package main

import (

	"fmt"
	"log"
	"os"

	"github.com/1amkaizen/telegoGPT/controllers"
	"github.com/1amkaizen/telegoGPT/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)






// Tambahkan variabel global untuk menyimpan log percakapan
var conversationLogs []controllers.ConversationLog

// Handler untuk menampilkan log percakapan dalam HTML
func logHandler(w http.ResponseWriter, r *http.Request) {
    // Gunakan template HTML untuk menampilkan log percakapan
    tmpl := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Log Percakapan</title>
    </head>
    <body>
        <h1>Log Percakapan</h1>
        <table>
            <tr>
                <th>Waktu</th>
                <th>Pengirim</th>
                <th>Pesan</th>
            </tr>
            {{range .}}
            <tr>
                <td>{{.Timestamp}}</td>
                <td>{{.Sender}}</td>
                <td>{{.Message}}</td>
            </tr>
            {{end}}
        </table>
    </body>
    </html>
    `

    // Buat instance template
    tmpl, err := template.New("log").Parse(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Kirim data log ke template dan tampilkan dalam HTML
    if err := tmpl.Execute(w, conversationLogs); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

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

    // Atur handler untuk menampilkan log
    http.HandleFunc("/log", logHandler)

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

		} else if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "twitter":
				// Tambahkan logika Anda di sini untuk menangani aksi yang dilakukan saat tombol diklik
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "https://twitter.com/1amkaizen")
				bot.Send(msg)
			case "github":
				// Tambahkan logika Anda di sini untuk menangani aksi yang dilakukan saat tombol diklik
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "https://github.com/1amkaizen/")
				bot.Send(msg)
			case "railway":
				// Tambahkan logika Anda di sini untuk menangani aksi yang dilakukan saat tombol diklik
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "https://railway.app?referralCode=v-jhtw")
				bot.Send(msg)
			case "replit":
				// Tambahkan logika Anda di sini untuk menangani aksi yang dilakukan saat tombol diklik
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "https://replit.com/@1amkaizen")
				bot.Send(msg)

			default:
				// Tambahkan logika Anda di sini untuk menangani aksi default jika tidak ada yang sesuai dengan callback.Data
			}
		}
}
