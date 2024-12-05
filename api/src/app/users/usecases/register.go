package usecases

import (
	"lib/app/users/models"
	"lib/common/errs"
	"lib/common/id"
	"lib/services"
)

type RegisterArgs struct {
	UserName models.UserName
	Password models.UserPassword
}

func Register(args RegisterArgs) any {
	db, freeDb := services.Db()
	defer freeDb()

	user, err := models.NewUser(
		id.New(),
		args.UserName,
		args.Password,
	)
	if err != nil {
		return err
	}

	var existing models.User

	result := db.Where(map[string]any{
		"user_name": user.UserName,
	}).First(&existing)

	if result.Error == nil {
		return errs.Exists("user with this user_name already exists")
	}

	result = db.Create(user)
	if result.Error != nil {
		return errs.Error("error registering %v", result.Error)
	}

	return nil
}
