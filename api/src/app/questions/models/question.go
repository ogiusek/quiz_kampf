package models

import (
	userModels "lib/app/users/models"
	"lib/common/id"
	"time"
)

type Question struct {
	QuestionId   id.ID           `gorm:"primaryKey"`
	CreatorId    id.ID           `gorm:"column:creator_id"`
	Creator      userModels.User `gorm:"foreignKey:user_id;references:creator_id"`
	QuestionText QuestionText    `gorm:"column:question_text"`
	Answer       Answer          `gorm:"column:answer"`
	CreatedAt    time.Time       `gorm:"column:created_at"`
}

func NewQuestion(id id.ID, creatorId id.ID, questionText QuestionText, answer Answer, createdAt time.Time) Question {
	return Question{
		QuestionId:   id,
		CreatorId:    creatorId,
		QuestionText: questionText,
		Answer:       answer,
		CreatedAt:    createdAt,
	}
}
