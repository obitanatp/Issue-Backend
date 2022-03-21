package handles

import (
	"goapp/pkg/db"
	"goapp/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func GetWord(c *fiber.Ctx) error {
	id := c.Params("id")
	var word models.Word
	result := db.Database.Preload("Answer").Find(&word, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&word)
}
