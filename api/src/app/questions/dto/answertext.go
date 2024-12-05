package dto

import "lib/app/questions/models"

type AnswerText struct {
	CorrectAnswers []models.AnswerMessage `json:"correct_answers"`
	CaseSensitive  bool                   `json:"case_sensitive"`
}

func AnswerTextDto(answer models.AnswerTextType) AnswerText {
	return AnswerText{
		CorrectAnswers: answer.CorrectAnswers,
		CaseSensitive:  answer.CaseSensitive,
	}
}
