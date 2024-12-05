package usecases

import (
	"lib/app/users/dto"
	"lib/app/users/models"
	"lib/services"
)

type ProfileArgs struct {
	Session dto.Session
}

func Profile(args ProfileArgs) any {
	db, freeDb := services.Db()
	defer freeDb()

	var user models.User
	db.Where(map[string]any{
		"user_id": args.Session.UserId,
	}).First(&user)

	return dto.ToUserDto(user)
}
