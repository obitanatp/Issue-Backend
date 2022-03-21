package main

import (
	"goapp/pkg/db"
	"goapp/pkg/handles"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	db.Connect()
	app.Get("/connect", handles.GetWords)
	app.Get("/connect/:id", handles.GetWord)
	app.Post("/connect", handles.AddWord2)
	app.Put("/connect/:id", handles.UpdateWord2)
	app.Delete("/connect/:id", handles.RemoveWord2)
	app.Post("/connectz", handles.AddWord3)
	app.Put("/connectz/:id", handles.UpdateWord3)
	app.Delete("/connectz/:id", handles.RemoveWord3)
	log.Fatal(app.Listen(":2000"))
}
