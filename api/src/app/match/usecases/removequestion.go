package usecases

import (
	"lib/app/match/messages"
	"lib/app/match/repo"
	userdto "lib/app/users/dto"
	"lib/common/id"
)

type RemoveQuestionArgs struct {
	Session    userdto.Session
	QuestionId id.ID `json:"question_id"`
}

func RemoveQuestion(args RemoveQuestionArgs) error {
	matchRepo := repo.GetMatchRepo()
	match, err := matchRepo.GetUserMatch(args.Session.UserId)
	if err != nil {
		return err
	}

	err = match.RemoveQuestion(args.QuestionId, args.Session.UserId)
	if err != nil {
		return err
	}

	err = matchRepo.UpdateMatch(match)
	if err != nil {
		return err
	}

	messages.SendMatch(match)

	return nil
}
