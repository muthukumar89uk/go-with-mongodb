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

	app.Post("/insert", h.InsertValue)
	app.Get("/get", h.GetValue)
	app.Put("/update/:id", h.UpdateById)
	app.Delete("/deleteDoc/:id", h.DeleteById)

	if err := app.Listen(":8080"); err != nil {
		fmt.Println("Connection Failed:", err.Error())
		return
	}
	
	fmt.Println("Server Connected !!!!!...")
}
