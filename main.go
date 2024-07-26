package main

import (
	"fmt"
	"sample/drivers"
	"sample/routers"
)

func main() {
	client, collection, err := drivers.DbConnection()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	routers.SetupRouter(client, collection)
}
