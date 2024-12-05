package usecases

import (
	"errors"
	"lib/app/questions/repo"
	userdto "lib/app/users/dto"
	"lib/common/errs"
	"lib/common/id"
)

var ErrIsNotAnOwner error = errors.New("you are not the owner of this question")
var ErrQuestionDoesNotExist error = errors.New("question does not exist")

type RemoveQuestionArgs struct {
	Session    userdto.Session ``
	QuestionId id.ID           `json:"question_id"`
}

func RemoveQuestion(args RemoveQuestionArgs) any {
	repo := repo.GetQuestionRepo()

	question := repo.GetQuestion(args.QuestionId)

	if question == nil {
		return errs.NotFound(ErrQuestionDoesNotExist.Error())
	}

	if question.CreatorId != args.Session.UserId {
		return errs.Forbidden(ErrIsNotAnOwner.Error())
	}

	if err := repo.RemoveQuestion(question.QuestionId); err != nil {
		return errs.NotFound(ErrQuestionDoesNotExist.Error())
	}

	return nil
}
