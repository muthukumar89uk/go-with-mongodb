package routers

import (
	"fmt"
	"sample/handlers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client, collection *mongo.Collection) {
	app := fiber.New()

	h := handlers.DataBase{
		Client:     client,
		Collection: collection,
	}

	app.Post("/create", h.InsertValue)
	app.Get("/getemployee", h.GetValue)
	app.Put("/update-employees/:id", h.UpdateById)
	app.Delete("/delete-employees/:id", h.DeleteById)

	if err := app.Listen(":8080"); err != nil {
		fmt.Println("Connection Failed:", err.Error())
		return
	}

	fmt.Println("Server Connected !!!!!...")
}
