package models

import (
	"lib/common/id"
	"lib/common/valid"
)

type User struct {
	UserId    id.ID     `gorm:"column:user_id;primaryKey"`
	UserImage UserImage `gorm:"column:user_image;null"`
	UserName  UserName  `gorm:"column:user_name;not null"`
	Hash      UserHash  `gorm:"column:hash;not null"`
}

func NewUser(id id.ID, userName UserName, password UserPassword) (*User, error) {
	user := User{
		UserId:   id,
		UserName: userName,
		Hash:     password.Hash(),
	}

	if err := valid.Valid(user); err != nil {
		return nil, err
	}

	return &user, nil
}
