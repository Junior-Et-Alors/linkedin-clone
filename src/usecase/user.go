package usecase

import (
	"context"
	"fmt"
	"linkedin-clone/db/mongodb"
	"linkedin-clone/src/entity"
	"linkedin-clone/src/repository/mongoRepository"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func InsertUser(user entity.User) {
	clientOptions := mongodb.Initializer()
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

	mongoRepository.InsertUser(context.Background(), userEntity, col)

}

func Login(user entity.User, passwordFromForm string) bool {
	clientOptions := mongodb.Initializer()
	u := mongoRepository.GetUser(clientOptions, user)

	match := bcrypt.CompareHashAndPassword(u.GetPassword(), []byte(user.GetMdp()))
	return match == nil
}

func CheckIsEmailExist(user entity.User) bool {
	clientOptions := mongodb.Initializer()
	u := mongoRepository.GetUser(clientOptions, user)

	return u.GetEmail() == user.GetEmail()
}
