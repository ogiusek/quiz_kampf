package usecases

import (
	"lib/app/match/messages"
	"lib/app/match/models"
	matchrepo "lib/app/match/repo"
	questionRepo "lib/app/questions/repo"
	userdto "lib/app/users/dto"
	"lib/common/id"
)

type AddQuestionArgs struct {
	Session    userdto.Session
	QuestionId id.ID `json:"question_id"`
}

func AddQuestion(args AddQuestionArgs) error {
	matchRepo := matchrepo.GetMatchRepo()
	match, err := matchRepo.GetUserMatch(args.Session.UserId)
	if err != nil {
		return err
	}

	questionRepo := questionRepo.GetQuestionRepo()
	question := questionRepo.GetQuestion(args.QuestionId)

	if question == nil {
		return models.ErrQuestionNotFound
	}

	err = match.AddQuestion(*question, args.Session.UserId)
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
