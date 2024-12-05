package usecases

import (
	"lib/app/questions/dto"
	questionsRepo "lib/app/questions/repo"
)

type SelectQuestionsArgs struct {
	SearchQuestion              string              `schema:"search_question"`
	SearchQuestionCaseSensitive bool                `schema:"search_question_case_sensitive"`
	SearchAnswer                string              `schema:"search_answer"`
	SearchAnswerCaseSensitive   bool                `schema:"search_answer_case_sensitive"`
	SearchNick                  string              `schema:"search_nick"`
	SearchNickCaseSentitive     bool                `schema:"search_nick_case_sensitive"`
	Limit                       questionsRepo.Limit `schema:"limit"`
	Page                        questionsRepo.Page  `schema:"page"`
}

func SelectQuestions(args SelectQuestionsArgs) any {
	repo := questionsRepo.GetQuestionRepo()
	questions := repo.Search(questionsRepo.SearchQuestionsArgs{
		SearchQuestion:              args.SearchQuestion,
		SearchQuestionCaseSensitive: args.SearchAnswerCaseSensitive,
		SearchAnswer:                args.SearchAnswer,
		SearchAnswerCaseSensitive:   args.SearchAnswerCaseSensitive,
		SearchNick:                  args.SearchNick,
		SearchNickCaseSentitive:     args.SearchAnswerCaseSensitive,
		Limit:                       args.Limit,
		Page:                        args.Page,
	})
	var result []dto.Question
	for _, question := range questions {
		result = append(result, dto.QuestionDto(question))
	}

	return result
}
