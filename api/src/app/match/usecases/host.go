package usecases

import (
	"lib/app/match/messages"
	"lib/app/match/models"
	matchrepo "lib/app/match/repo"
	userdto "lib/app/users/dto"
	"lib/app/users/repo"
)

type HostArgs struct {
	Session userdto.Session
}

func Host(args HostArgs) error {
	matchRepo := matchrepo.GetMatchRepo()
	_, err := matchRepo.GetUserMatch(args.Session.UserId)
	if err == nil {
		return ErrAlreadyInMatch
	}

	userRepo := repo.GetUserRepo()
	user := userRepo.GetById(args.Session.UserId)
	if user == nil {
		return models.ErrUserNotFound
	}
	player := models.NewAdmin(*user)
	match, err := models.NewMatch(player)
	if err != nil {
		return err
	}

	err = matchRepo.AddMatch(match)
	if err != nil {
		return err
	}

	go messages.SendMatch(match)

	return nil
}
