package Database

import (
	"context"
	"fmt"

	"g_o/MongoDb/controllers"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getDatabaseUri() string {
	return "mongodb://127.0.0.1:27017/Users"
}

func ConnectMongo() {
	fmt.Println("Connecting to mongo...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(getDatabaseUri())
	DB, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	err = DB.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Failed to ping MongoDB:", err)
		return
	}
	database := DB.Database("Users")
	userCollection := database.Collection("users")
	controllers.InitializeUserDatabase(userCollection)
	fmt.Println("Successfully connected to MongoDB!")
}
