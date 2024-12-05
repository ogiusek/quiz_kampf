package dto

import (
	"lib/app/users/models"
	"lib/common/id"
)

type User struct {
	Id        id.ID            `json:"id"`
	UserName  models.UserName  `json:"user_name"`
	UserImage models.UserImage `json:"user_image"`
}

func ToUserDto(user models.User) User {
	return User{
		Id:        user.UserId,
		UserName:  user.UserName,
		UserImage: user.UserImage,
	}
}
