package models

type User struct {
	Name        string `json:"name" bson:"name"`
	Address     string `json:"address" bson:"address"`
	Date        string `json:"date" bson:"date"`
	PhoneNumber string `json:"phonenumber" bson:"phonenumber"`
}
