package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	Database "g_o/MongoDb/Database"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

// InitializeUserDatabase initializes the MongoDB collection.
func InitializeUserDatabase(Collection *mongo.Collection) {
	userCollection = Collection
}

// CreateUserHandler creates a new user in the database.
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user Database.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created Successfully"})
}

// Placeholder handlers for other CRUD operations
func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {}
func GetUserHandler(w http.ResponseWriter, r *http.Request)    {}
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {}
