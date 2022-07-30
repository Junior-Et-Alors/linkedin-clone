package usecase

import (
	"context"
	"fmt"
	"linkedin-clone/src/entity"
	"linkedin-clone/src/repository"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func InsertUser(clientOptions *options.ClientOptions, user entity.User) {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo connection error :", err)
	}

	col := client.Database("linkedin-clone").Collection("user")

	userEntity := entity.User{
		Email:     user.GetEmail(),
		Password:  user.GetPassword(),
		Firstname: user.GetFirstName(),
		Lastname:  user.GetLastname(),
		Age:       user.GetAge(),
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	repository.InsertUser(context.Background(), userEntity, col)

}

func Login(clientOptions *options.ClientOptions, user entity.User, passwordFromForm string) bool {

	u := repository.GetUser(clientOptions, user)

	match := bcrypt.CompareHashAndPassword(u.GetPassword(), []byte(user.GetMdp()))
	return match == nil
}

func CheckIsEmailExist(clientOptions *options.ClientOptions, user entity.User) bool {
	u := repository.GetUser(clientOptions, user)
	fmt.Println(u.GetEmail(), user.GetEmail())
	return u.GetEmail() == user.GetEmail()
}
