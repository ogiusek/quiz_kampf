package usecases

import (
	"lib/app/users/dto"
	"lib/app/users/models"
	"lib/common/errs"
	"lib/services"
)

type RefreshArgs struct {
	SessionToken dto.SessionToken `json:"session_token"`
	RefreshToken dto.RefreshToken `json:"refresh_token"`
}

func Refresh(args RefreshArgs) any {
	// _, err := args.SessionToken.DecodeSession()
	// if err == nil {
	// 	return errs.BadRequest("session token has to expire")
	// }

	refreshPayload, err := args.RefreshToken.DecodeSession()
	if err != nil {
		return errs.BadRequest(err.Error())
	}

	if !args.SessionToken.HashMatchesSessionToken(refreshPayload.Hash) {
		return errs.BadRequest("invalid session token")
	}

	db, freeDb := services.Db()
	defer freeDb()

	var user models.User

	result := db.Where(map[string]any{
		"user_id": refreshPayload.UserId,
	}).First(&user)

	if result.Error != nil {
		return errs.NotFound("User do not exists")
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

	return dto.Tokens{ // this uses login response
		SessionToken: sessionToken,
		RefreshToken: refreshToken,
	}
}
