package handles

import (
	"goapp/pkg/db"
	"goapp/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func GetWords(c *fiber.Ctx) error {
	var words []models.Word
	db.Database.Preload("Answer").Order("id").Find(&words)
	return c.Status(200).JSON(words)
}
