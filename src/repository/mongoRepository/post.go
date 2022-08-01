package mongoRepository

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoFields struct {
	FieldStr  string
	FieldInt  int
	FieldBool bool
}

func Collection() {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo connection error :", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("linkedin-clone").Collection("post")

	oneDoc := MongoFields{
		FieldStr:  "exemple teste",
		FieldInt:  1234,
		FieldBool: true,
	}

	fmt.Println("oneDoc type : ", reflect.TypeOf(oneDoc))

	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("error insert : ", insertErr)
		os.Exit(1)
	} else {
		fmt.Println("insert result type : ", reflect.TypeOf(result))
		fmt.Println("insert ok : ", result)

		newId := result.InsertedID
		fmt.Println("new id type : ", reflect.TypeOf(newId))
		fmt.Println("new id : ", newId)
	}

}
