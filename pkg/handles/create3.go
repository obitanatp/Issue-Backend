package handles

import (
	"goapp/pkg/db"
	"goapp/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func AddWord3(c *fiber.Ctx) error {
	ans := new(models.WordAnswer)
	if err := c.BodyParser(ans); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Database.Create(&ans)
	return c.Status(200).JSON(ans)
}
