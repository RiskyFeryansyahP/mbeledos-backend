package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/confus1on/mbeledos/models"
	"github.com/confus1on/mbeledos/user"
)

type MongoUserRepository struct {
	DB *mongo.Database
}

func NewMongoUserRepository(db *mongo.Database) user.Repository {
	return &MongoUserRepository{DB: db}
}

func (ur *MongoUserRepository) Login(phonenumber string) (res *models.User, err error) {
	userCollection := ur.DB.Collection("users")

	filter := bson.D{
		primitive.E{
			Key:   "phonenumber",
			Value: phonenumber,
		},
	}

	err = userCollection.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ur *MongoUserRepository) Register(user *models.User) error {
	userCollection := ur.DB.Collection("users")

	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	if result.InsertedID == nil {
		log.Println("mongo_user : Failed to insert data", err.Error())
	}

	return nil

}
