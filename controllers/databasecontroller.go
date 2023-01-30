package controllers

import (
	"net/http"
	"telegoGPT/models"

	"github.com/gofiber/fiber/v2"
)

func GetDataFromTelegramBot(c *fiber.Ctx) error {
	// Ambil data dari database di Telegram Bot project
	var data []models.Users
	models.DB.Find(&data)

	if err := models.DB.Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving data from Telegram Bot database",
		})
	}

	// Kembalikan data dalam bentuk JSON
	return c.JSON(data)
}

func getDataFromTelegramBotDB() (interface{}, error) {
	// Logic untuk mengambil data dari database
	// Sudah digantikan oleh baris pada function GetDataFromTelegramBot
	return nil, nil
}
