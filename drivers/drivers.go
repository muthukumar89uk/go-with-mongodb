package drivers

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DbConnection() (*mongo.Client, *mongo.Collection, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("DataBase Not Connected", err)
		return nil, nil, err
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println("Check DataBase Connected or Not", err)
		return nil, nil, err
	}

	fmt.Println("DataBase Connected....")

	collection := client.Database("testDB").Collection("employee")

	return client, collection, nil
}
