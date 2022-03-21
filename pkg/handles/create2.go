package handles

import (
	"goapp/pkg/db"
	"goapp/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func AddWord2(c *fiber.Ctx) error {
	ans := new(models.Word)
	if err := c.BodyParser(ans); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Database.Preload("Answers").Create(&ans)
	return c.Status(200).JSON(ans)
}
