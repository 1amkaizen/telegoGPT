package controllers

import (
	"net/http"

	"github.com/1amkaizen/telegoGPT/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.Users
	err := models.DB.Find(&users).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while fetching users",
		})
	}
	return c.JSON(users)
}
