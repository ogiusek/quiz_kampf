package dto

import "lib/app/questions/models"

type AnswerOptions struct {
	Answers []models.AnswerMessage `json:"answers"`
	Correct int                    `json:"correct"`
}

func AnswerOptionsDto(answer models.AnswerOptionsType) AnswerOptions {
	return AnswerOptions{
		Answers: answer.Answers,
		Correct: answer.Correct,
	}
}
