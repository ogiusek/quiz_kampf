package usecases

import (
	"lib/app/match/messages"
	"lib/app/match/models"
	matchrepo "lib/app/match/repo"
	userdto "lib/app/users/dto"
	"lib/common/id"
)

type SetAnswerTimeArgs struct {
	Session             userdto.Session
	QuestionId          id.ID                      `json:"question_id"`
	AnswerTimeInSeconds models.AnswerTimeInSeconds `json:"answer_time_in_seconds"`
}

func SetAnswerTime(args SetAnswerTimeArgs) error {
	matchRepo := matchrepo.GetMatchRepo()
	match, err := matchRepo.GetUserMatch(args.Session.UserId)
	if err != nil {
		return err
	}

	for i := range match.Questions {
		if match.Questions[i].QuestionId == args.QuestionId {
			match.Questions[i].SetAnswerTime(args.AnswerTimeInSeconds)
			err := matchRepo.UpdateMatch(match)
			if err != nil {
				return err
			}
			messages.SendMatch(match)
			return nil
		}
	}

	return models.ErrQuestionNotFound
}
