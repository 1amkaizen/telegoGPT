package controllers

import (
	"net/http"

	"github.com/1amkaizen/telegoGPT/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetDataFromTelegramBot(c *fiber.Ctx) error {
	id := c.Params("id")
	var data []models.Users
	models.DB.Find(&data)
	if err := models.DB.First(&data, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Render("index", fiber.Map{
				"alert": "Data not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.Render("index", fiber.Map{
		"data": data,
	})
}
