package usecases

import (
	"lib/app/questions/dto"
	"lib/app/questions/models"
	"lib/app/questions/repo"
	usersdto "lib/app/users/dto"
	"lib/common/errs"
	"lib/common/id"
	"time"
)

type AddQuestionArgs struct {
	Session      usersdto.Session
	QuestionText models.QuestionText `json:"question"`
	Answer       models.RawAnswer    `json:"answer"`
}

func AddQuestion(args AddQuestionArgs) any {
	repo := repo.GetQuestionRepo()
	answer, err := args.Answer.ToAnswer()

	if err != nil {
		return errs.BadRequest(err.Error())
	}

	question := models.NewQuestion(
		id.New(),
		args.Session.UserId,
		args.QuestionText,
		answer,
		time.Now(),
	)

	if err := repo.AddQuestion(question); err != nil {
		return errs.Error("error creating question: %v", err)
	}

	return dto.QuestionDto(question)
}
