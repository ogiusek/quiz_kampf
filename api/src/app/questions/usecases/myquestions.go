package usecases

import (
	"lib/app/questions/dto"
	questionsrepo "lib/app/questions/repo"
	userdto "lib/app/users/dto"
)

type MyQuestionsArgs struct {
	session userdto.Session    `schema:""`
	page    questionsrepo.Page `schema:"page"`
}

func MyQuestions(args MyQuestionsArgs) any {
	repo := questionsrepo.GetQuestionRepo()
	questions := repo.UserQuestions(args.session.UserId, args.page)
	var result []dto.Question
	for _, question := range questions {
		result = append(result, dto.QuestionDto(question))
	}
	return result
}
