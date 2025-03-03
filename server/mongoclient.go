package main 

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
  
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGO_DBNAME = os.Getenv("MONGO_DBNAME") 
	MONGO_PASSWORD = os.Getenv("MONGO_PASSWORD") 
	MONGO_USERNAME = os.Getenv("MONGO_USERNAME") 
	MONGO_URI = fmt.Sprintf(os.Getenv("MONGO_URI"), MONGO_USERNAME, MONGO_PASSWORD, MONGO_DBNAME)
) 
  
func NewMongoClient() (*mongo.Client, error) {

	// load env file variables 
	err := godotenv.Load(); 
	if err != nil {
		return nil, err
	} 

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGO_URI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err 
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil { 
			panic(err) 
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil { 
		return nil, err 
	}

	log.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client, nil
}
