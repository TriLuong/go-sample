package database

import (
	"context"
	"fmt"
	"log"

	"github.com/TriLuong/go-sample/models"
	"go.mongodb.org/mongo-driver/bson"
)

type Model struct {
	ModelName string
}

func (m Model) AddNew(user interface{}) error {
	client := GetMongoClient()
	collection := client.Database("go-sample").Collection(m.ModelName)
	_, err := collection.InsertOne(context.TODO(), user)

	return err
}

func (m Model) GetUsers() ([]models.User, error) {
	client := GetMongoClient()
	collection := client.Database("go-sample").Collection(m.ModelName)
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println("Error")
	}

	var users []models.User

	for cur.Next(context.TODO()) {
		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		users = append(users, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return users, nil

}
