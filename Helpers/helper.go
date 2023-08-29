package helpers

import (
	"context"
	"fmt"
	"log"

	model "github.com/sachinchaudhary003/golangAuth/Model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoconn = "mongodb+srv://Godb:Golang123@cluster0.0r5zajl.mongodb.net/"
const dbName = "Auhtentication"
const colName = "auth"

var Collection *mongo.Collection

func init() {

	clientoption := options.Client().ApplyURI(mongoconn)

	client, err := mongo.Connect(context.TODO(), clientoption)

	if err != nil {
		log.Fatal(err)
	}
	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("Connected to mongo")

}

func InsertOne(user model.User) {
	inserted, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("iserted user ", inserted.InsertedID)
}
