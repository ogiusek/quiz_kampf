package usecases

import (
	"lib/app/match/messages"
	matchrepo "lib/app/match/repo"
	userdto "lib/app/users/dto"
)

type StartArgs struct {
	Session userdto.Session
}

func Start(args StartArgs) error {
	matchRepo := matchrepo.GetMatchRepo()
	match, err := matchRepo.GetUserMatch(args.Session.UserId)
	if err != nil {
		return err
	}

	if err := match.CanStart(); err != nil {
		return err
	}
	messages.SendStarted(match)
	match.Start()
	err = matchRepo.UpdateMatch(match)
	if err != nil {
		return nil
	}

	messages.SendQuestion(match)
	EnsureAnswered(EnsureAnswredArgs{Match: match})

	return nil
}
