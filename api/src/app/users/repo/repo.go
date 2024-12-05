package repo

import (
	"lib/app/users/models"
	"lib/common/id"
	"lib/services"
)

type UserRepo interface {
	GetById(id.ID) *models.User
}

type userRepo struct{}

func (userRepo) GetById(userId id.ID) *models.User {
	db, free := services.Db()
	defer free()
	var user models.User
	if err := db.Where("user_id = ?", userId).Find(&user).Error; err != nil {
		return nil
	}

	return &user
}

func GetUserRepo() UserRepo {
	return &userRepo{}
}
