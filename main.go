package main

import (
	"net/http"
    "encoding/json"
	
	"fmt"
	"log"
	"os"

	"github.com/1amkaizen/telegoGPT/controllers"
	"github.com/1amkaizen/telegoGPT/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


// Endpoint API untuk mengambil data pengguna dan percakapan
func UserConversationHandler(w http.ResponseWriter, r *http.Request) {
    // Di sini, Anda perlu mengambil data pengguna dan percakapan dari database Anda.
    // Gantilah kode berikut ini dengan logika yang sesuai untuk mengambil data tersebut.
    userData := []models.UserData{}      // Ganti dengan model data pengguna Anda
    conversationData := []models.Message{} // Ganti dengan model data percakapan Anda

    // Menggabungkan data pengguna dan percakapan ke dalam struktur data yang sesuai
    data := struct {
        Users       []models.UserData
        Conversations []models.Message
    }{
        Users:       userData,
        Conversations: conversationData,
    }

    // Mengubah data ke dalam format JSON
    jsonData, err := json.Marshal(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Menetapkan tipe konten dan mengirimkan respons JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
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


// Konfigurasi lainnya ...

    // Mengatur endpoint API untuk mengambil data pengguna dan percakapan
    http.HandleFunc("/api/user-conversation", UserConversationHandler)

    // Mulai server HTTP
    go func() {
        log.Fatal(http.ListenAndServe(":8080", nil)) // Ganti port sesuai kebutuhan Anda
    }()

    // Logika bot Anda ...

	
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
	
	

	
}



