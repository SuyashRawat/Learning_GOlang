package main

import (
	"fmt"
	"go-crud/packages/constants"
	"go-crud/packages/db"
	"go-crud/packages/handlers"
	"go-crud/packages/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable dbname=%s", constants.DBHost, constants.DBPort, constants.DBUser, constants.DBPass, constants.DBDbase)
	fmt.Println("Database URL: ", dbURL)

	// Connection to Database
	database, err := db.InitDB(dbURL)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	defer database.Close()

	// Create the table if it doesn't exist
	_, _ = database.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")

	defer database.Close() // Closing the database

	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.GetUsers(database)).Methods("GET")            // Get all users
	router.HandleFunc("/users/{id}", handlers.GetUserByID(database)).Methods("GET")    // Get user by id
	router.HandleFunc("/users", handlers.CreateUser(database)).Methods("POST")         // Create user
	router.HandleFunc("/users/{id}", handlers.UpdateUserByID(database)).Methods("PUT") // Update user by id
	router.HandleFunc("/users/{id}", handlers.DeleteUser(database)).Methods("DELETE")

	// log.Fatal(http.ListenAndServe(":8080", router))
	// Instead of using http.ListenAndServe(":8080", router) we will add a middleware function
	log.Fatal(http.ListenAndServe(":8080", middlewares.JsonContentTypeMiddleware(router)))
}

// We are using a middleware function to set the content type of the response to application/json.
// This middleware function will be called before the actual handler function is called.
// Hence htis header will be added to all the resposnes by default
// If we do not wish to achieve this, we will have to set the header individually in each handler function.
