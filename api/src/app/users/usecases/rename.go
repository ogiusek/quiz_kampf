package usecases

import (
	"errors"
	"lib/app/users/dto"
	"lib/app/users/models"
	"lib/common/errs"
	"lib/services"
)

type RenameArgs struct {
	Session dto.Session
	NewName models.UserName `json:"new_name"`
}

var ErrNickIsAlreadyTaken error = errors.New("nick is already taken")

func Rename(args RenameArgs) any {
	if err := args.Session.HasRight(dto.Permission("rename")); err != nil {
		return err
	}

	db, freeDb := services.Db()
	defer freeDb()

	var user models.User

	err := db.Model(&models.User{}).Where("user_name = ?", args.NewName).First(&user).Error

	if err == nil {
		return errs.Conflict(ErrNickIsAlreadyTaken.Error())
	}

	db.Model(&models.User{}).Where(map[string]any{
		"user_id": args.Session.UserId,
	}).Updates(map[string]interface{}{
		"user_name": args.NewName,
	})

	return nil
}
