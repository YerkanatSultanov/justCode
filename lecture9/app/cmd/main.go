package main

import (
	"fmt"
	"justCode/lecture9/app/model"
	"justCode/lecture9/app/repository/followers"
	"justCode/lecture9/app/repository/image"
	"justCode/lecture9/app/repository/user"
)

func main() {
	//IMPORTANT: Просто как пример сохраняю все в массивах, в проекте буду работать с базами данных

	userRepository := user.NewUserRepository()
	imageRepository := image.NewImageRepository()
	followersRepository := followers.NewFollowersRepository()

	user1 := &model.User{
		Id:       1,
		Username: "john_doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	err := userRepository.Create(user1)
	if err != nil {
		return
	}

	image1 := &model.Image{
		Image_id:    1,
		Image_link:  "https://example.com/image1.jpg",
		Description: "Example image 1",
		UserID:      1,
	}
	err = imageRepository.Create(image1)
	if err != nil {
		return
	}

	follower1 := &model.Followers{
		UserID:      1,
		Follower_id: 2,
	}
	err = followersRepository.Create(follower1)
	if err != nil {
		return
	}

	allUsers := userRepository.GetAll()
	allImages := imageRepository.GetAll()
	allFollowers := followersRepository.GetAll()

	fmt.Println("All Users:", allUsers)
	fmt.Println("All Images:", allImages)
	fmt.Println("All Followers:", allFollowers)
}
