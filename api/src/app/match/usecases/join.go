package usecases

import (
	"lib/app/match/messages"
	matchmodels "lib/app/match/models"
	matchrepo "lib/app/match/repo"
	userdto "lib/app/users/dto"
	"lib/app/users/repo"
	"lib/common/id"
)

type JoinAgrs struct {
	Session userdto.Session // this is injected
	MatchId id.ID           `json:"match_id"`
}

func Join(args JoinAgrs) error {
	matchRepo := matchrepo.GetMatchRepo()
	_, err := matchRepo.GetUserMatch(args.Session.UserId)
	if err == nil {
		return ErrAlreadyInMatch
	}

	match, err := matchRepo.GetMatch(args.MatchId)
	if err != nil {
		return err
	}
	userRepo := repo.GetUserRepo()
	user := userRepo.GetById(args.Session.UserId)
	if user == nil {
		return matchmodels.ErrUserNotFound
	}

	player := matchmodels.NewPlayer(*user)
	err = match.Join(player)
	if err != nil {
		return err
	}

	matchRepo.UpdateMatch(match)

	go messages.SendMatch(match)

	return nil
}
