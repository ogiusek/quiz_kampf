package usecases

import (
	"lib/app/questions/models"
	"lib/app/questions/repo"
	usersdto "lib/app/users/dto"
	"lib/common/errs"
	"lib/common/id"
)

type UpdateQuestionArgs struct {
	Session      usersdto.Session
	QuestionId   id.ID               `json:"question_id"`
	QuestionText models.QuestionText `json:"question"`
	Answer       models.RawAnswer    `json:"answer"`
}

func UpdateQuestion(args UpdateQuestionArgs) any {
	answer, err := args.Answer.ToAnswer()
	if err != nil {
		return errs.BadRequest("wrong answer")
	}

	repo := repo.GetQuestionRepo()
	question := repo.GetQuestion(args.QuestionId)

	if question == nil {
		return errs.NotFound(ErrQuestionDoesNotExist.Error())
	}

	if question.CreatorId == args.Session.UserId {
		return errs.Forbidden(ErrIsNotAnOwner.Error())
	}

	question.QuestionText = args.QuestionText
	question.Answer = answer

	if err := repo.UpdateQuestion(*question); err != nil {
		return errs.NotFound(ErrQuestionDoesNotExist.Error())
	}

	return nil
}
