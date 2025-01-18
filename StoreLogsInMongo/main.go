package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	Mongomodels "startlearing/Schema"
	config "startlearing/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func create(w http.ResponseWriter, r *http.Request, i int) {
	userLogs := &Mongomodels.UserLogs{
		// ServerID:  "server123", // Example value
		Goroutine: i,
		Timestamp: time.Now(),
		Message:   "Fake Error" + strconv.Itoa(i),
		// Page:      "/example-page", // Example value
	}

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the log into the MongoDB collection
	_, err := userCollection.InsertOne(ctx, userLogs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "UserLog created successfully"})
}

func CreateLogs(w http.ResponseWriter, r *http.Request) {
	// Create user logs
	for i := 0; i < 5; i++ {
		go create(w, r, i+1)
	}

}

func main() {
	// Connect to MongoDB
	config.ConnectToMongo()
	if config.DB == nil {
		log.Fatal("Failed to connect to MongoDB.")
		return
	}

	dbase := config.DB.Database("UsersLogs")
	userCollection = dbase.Collection("UsersLogs")
	http.HandleFunc("/create-log", CreateLogs)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
