# Go MongoDB CRUD Example

This repository demonstrates how to set up a Go project to perform CRUD (Create, Read, Update, Delete) operations using MongoDB.

## Features

- MongoDB connection and operations in Go.
- CRUD operations.

## Requirements

- Go 1.15 or higher
- MongoDB
- Fiber Go Web Framework

## Getting Started

### Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/muthukumar89uk/go-with-mongodb.git
    ```
 Click here to directly [download it](https://github.com/muthukumar89uk/go-with-mongodb/zipball/master).

2. **Install Go dependencies:**

    ```sh
    go mod tidy
    ```

### Setup MongoDB

1. **Locally:**

    Install MongoDB from the [official MongoDB website](https://www.mongodb.com/try/download/community).

### Go Application

1. **Create the `Employee` struct:**

    Create a `models` directory and an `models.go` file with the following content:

    ```go
    type Employee struct {
	    Id   string `json:"id,omitempty"       bson:"_id,omitempty"`
	    Name string `json:"name,omitempty"     bson:"name,omitempty"`
	    Age  int    `json:"age,omitempty"      bson:"age,omitempty"`
    }
    ```

### Run the Application

1. **Run the Go application:**

    ```sh
    go run .
    ```

2. **API Endpoints:**

    - Create an employee: `POST /create`
    - Get an employee by ID: `GET /getemployee`
    - Update an employee by ID: `PUT /update-employees/:id`
    - Delete an employee by ID: `DELETE /delete-employees/:id`



