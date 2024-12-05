package usecases

import (
	"lib/app/match/messages"
	"lib/app/match/repo"
	userdto "lib/app/users/dto"
)

type QuitArgs struct {
	Session userdto.Session
}

func Quit(args QuitArgs) error {
	matchRepo := repo.GetMatchRepo()
	match, err := matchRepo.GetUserMatch(args.Session.UserId)
	if err != nil {
		return err
	}

	err = match.Quit(args.Session.UserId)
	if err != nil {
		return err
	}

	if len(match.Players) == 0 {
		err = matchRepo.RemoveMatch(match.MatchId)
	} else {
		err = matchRepo.UpdateMatch(match)
	}

	if err != nil {
		return err
	}

	messages.SendMatch(match)

	return nil
}
