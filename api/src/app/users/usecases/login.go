package usecases

import (
	"errors"
	"lib/app/users/dto"
	"lib/app/users/models"
	"lib/common/errs"
	"lib/services"

	"gorm.io/gorm"
)

type LoginArgs struct {
	UserName models.UserName
	Password models.UserPassword
}

func Login(args LoginArgs) any {
	db, freeDb := services.Db()
	defer freeDb()
	var user models.User

	result := db.Where(map[string]any{
		"user_name": args.UserName,
	}).First(&user)

	// this 2 can have the same error message for robust security
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errs.NotFound("user with this name does not exist")
		}
		return errs.Error("error logging in %v", result.Error)
	}

	if !args.Password.Matches(user.Hash) {
		return errs.InvalidInput("invalid password")
	}

	session := dto.ToSession(user)
	sessionToken, err := session.Encode()
	if err != nil {
		return err
	}

	refreshToken, err := sessionToken.RefreshToken(user.UserId)
	if err != nil {
		return err
	}

	return dto.Tokens{
		SessionToken: sessionToken,
		RefreshToken: refreshToken,
	}
}
