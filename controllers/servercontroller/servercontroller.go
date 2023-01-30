package servercontroller

import (
	"net/http"
	"telegoGPT/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var users []models.Users
	models.DB.Find(&users)
	//      return c.Status(fiber.StatusOK).JSON(books)
	return c.Render("index", fiber.Map{
		"Title": "world",
	})
}

func Show(c *fiber.Ctx) error {
	id := c.Params("id")
	var users []models.Users
	models.DB.Find(&users)
	if err := models.DB.First(&users, id).Error; err != nil {
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
		"user": users,
	})
}

func Create(c *fiber.Ctx) error {
	var users models.Users
	if err := c.BodyParser(&users); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Render("index", fiber.Map{
		"alert": "Data berhasil ditambahkan",
	})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var users models.Users

	if err := c.BodyParser(&users); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&users).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tidak dapat mengupdate data",
		})

	}
	return c.JSON(fiber.Map{
		"message": "data berhasil di update",
	})

}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var users models.Users

	if models.DB.Delete(&users, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak dapat menghapus data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil di hapus",
	})
}
