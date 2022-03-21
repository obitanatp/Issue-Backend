package handles

import (
	"fmt"
	"goapp/pkg/db"
	"goapp/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func UpdateWord2(c *fiber.Ctx) error {
	word := new(models.Word)
	id := c.Params("id")

	if err := c.BodyParser(word); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	fmt.Println(word)
	db.Database.Preload("Answers").Where("id = ?", id).Updates(&word)
	return c.Status(200).JSON(word)
}
