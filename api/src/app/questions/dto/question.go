package dto

import (
	"lib/app/questions/models"
	usersDto "lib/app/users/dto"
	"lib/common/id"
	"time"
)

type Question struct {
	// CreatorId    id.ID               `json:"creator_id"`
	QuestionId   id.ID               `json:"question_id"`
	CreatorId    id.ID               `json:"-"`
	Creator      usersDto.User       `json:"creator"`
	QuestionText models.QuestionText `json:"question"`
	Answer       Answer              `json:"answer"`
	CreatedAt    time.Time           `json:"created_at"`
}

func QuestionDto(model models.Question) Question {
	return Question{
		QuestionId:   model.QuestionId,
		CreatorId:    model.CreatorId,
		Creator:      usersDto.ToUserDto(model.Creator),
		QuestionText: model.QuestionText,
		Answer:       AnswerDto(model.Answer),
		CreatedAt:    model.CreatedAt,
	}
}
