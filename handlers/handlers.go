package handlers

import (
	"sample/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataBase struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func (d DataBase) InsertValue(c *fiber.Ctx) error {
	employee := new(models.Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	// mongoDB automatically created ID so initial is empty
	employee.Id = ""

	result, err := d.Collection.InsertOne(c.Context(), employee)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to Insert the document",
			"error":   err.Error(),
		})
	}

	filter := bson.D{{Key: "_id", Value: result.InsertedID}}

	// ReCheck in DB the value is Inserted or not
	createRecord := d.Collection.FindOne(c.Context(), filter)

	createRecord.Decode(employee)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    employee,
		"message": "Employee document inserted successfully",
	})
}

func (d DataBase) GetValue(c *fiber.Ctx) error {
	query := bson.D{{}}
	cursor, err := d.Collection.Find(c.Context(), query)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to Insert the document",
			"error":   err.Error(),
		})
	}

	var employees []models.Employee = make([]models.Employee, 0)

	if err = cursor.All(c.Context(), &employees); err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to get the document",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    employees,
		"message": "Successfully Fetched All employee document",
	})
}

func (d DataBase) UpdateById(c *fiber.Ctx) error {
	idParam := c.Params("id")

	employeeId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Not a valid ObjectId",
			"error":   err.Error(),
		})
	}

	var employee = new(models.Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to parse the body",
			"error":   err.Error(),
		})
	}

	query := bson.D{{Key: "_id", Value: employeeId}}

	update := bson.D{
		{
			Key:   "$set",
			Value: employee,
		},
	}

	_, err = d.Collection.UpdateOne(c.Context(), query, update)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to update the document",
			"error":   err.Error(),
		})
	}

	updatedEmp := new(models.Employee)

	cursor := d.Collection.FindOne(c.Context(), query)

	err = cursor.Decode(updatedEmp)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to update the document",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    updatedEmp,
		"message": "Successfully Updated the given Employee Id's document",
	})
}

func (d DataBase) DeleteById(c *fiber.Ctx) error {
	empId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Not a valid ObjectId",
			"error":   err.Error(),
		})
	}

	query := bson.D{{Key: "_id", Value: empId}}

	result, err := d.Collection.DeleteOne(c.Context(), &query)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to Delete the document",
			"error":   err.Error(),
		})
	}

	if result.DeletedCount < 1 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Already this Document deleted",
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Successfully Deleted the document",
	})
}
