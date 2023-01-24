package config

import (
	"UserAuth/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"

var Client *mongo.Client

func init() {
	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	utils.CheckNilErr(err)
	fmt.Println("MongoDB connection success")
}
