package dto

import (
	"fmt"
	"lib/app/questions/models"
)

type Answer struct {
	AnswerType models.AnswerType `json:"answer_type"`
	AnswerData any               `json:"answer_data"`
}

func AnswerDto(model models.Answer) Answer {
	var answerData any
	switch v := model.AnswerData.(type) {
	case models.AnswerOptionsType:
		answerData = AnswerOptionsDto(v)
	case models.AnswerTextType:
		answerData = AnswerTextDto(v)
	default:
		panic(fmt.Sprintf("application do not has dto for answer of type %T", answerData))
	}
	return Answer{
		AnswerType: models.AnswerOptions,
		AnswerData: answerData,
	}
}
