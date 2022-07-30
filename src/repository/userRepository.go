package repository

import (
	"context"
	"fmt"
	"linkedin-clone/src/entity"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertUser(ctx context.Context, user entity.User, col *mongo.Collection) *mongo.InsertOneResult {

	result, _ := col.InsertOne(ctx, user)
	return result
	// if insertErr != nil {
	// 	fmt.Println("error insert : ", insertErr)
	// 	os.Exit(1)
	// } else {

	// 	newId := result.InsertedID
	// 	fmt.Println("new id type : ", reflect.TypeOf(newId))
	// 	fmt.Println("new id : ", newId)
	// }

}

func GetUser(clientOptions *options.ClientOptions, user entity.User) entity.User {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo connection error :", err)
	}

	collection := client.Database("linkedin-clone").Collection("user")
	filter := bson.D{{Key: "email", Value: user.GetEmail()}}

	insert := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}
	if insert != nil {
		user.SetEmail("")
		return user
	}

	return user
}
