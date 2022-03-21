package handles

import (
	"goapp/pkg/db"
	"goapp/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func RemoveWord2(c *fiber.Ctx) error {
	id := c.Params("id")
	var word models.Word

	result := db.Database.Preload("Answers").Delete(&word, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
