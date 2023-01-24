package repository

import (
	"UserAuth/config"
	"UserAuth/entities"
	"UserAuth/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type UserRepository struct {
}

var collection = config.Client.Database("go_auth").Collection("users")

func (UserRepository) CreateUser(user *entities.User) primitive.ObjectID {
	inserted, err := collection.InsertOne(context.Background(), user)
	utils.CheckNilErr(err)
	return inserted.InsertedID.(primitive.ObjectID)
}

func (UserRepository) GetUserByUserName(username string) *entities.User {
	filter := bson.D{{"userName", username}}
	var user entities.User
	cur, err := collection.Find(context.Background(), filter)
	utils.CheckNilErr(err)
	defer cur.Close(context.Background())

	if cur.TryNext(context.Background()) {
		cur.Decode(&user)
		return &user
	} else {
		return new(entities.User)
	}

}

func (UserRepository) GetAllUsers() []entities.User {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var users []entities.User

	for cur.Next(context.Background()) {
		var user entities.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer cur.Close(context.Background())
	return users
}
