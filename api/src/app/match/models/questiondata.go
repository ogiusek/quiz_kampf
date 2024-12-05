package models

import (
	"errors"
	"lib/app/questions/models"
	"lib/common/id"
	"time"
)

var ErrAlreadyAnswered error = errors.New("question got already answered")

type QuestionData struct {
	QuestionId          id.ID               `json:"question_id"`
	Question            models.Question     `json:"question"`
	AnswerTimeInSeconds AnswerTimeInSeconds `json:"answer_time_in_seconds"`
	Started             *time.Time          `json:"started"`
	AnsweredBy          *id.ID              `json:"answered_by"`
	Answered            *time.Time          `json:"answered"`
	AnsweredCorrectly   *bool               `json:"correct_answer"`
}

func NewQuestion(question models.Question) QuestionData {
	return QuestionData{
		QuestionId:          question.QuestionId,
		Question:            question,
		AnswerTimeInSeconds: DefaultAnswerTime,
	}
}

func (answerData *QuestionData) SetAnswerTime(answerTime AnswerTimeInSeconds) {
	answerData.AnswerTimeInSeconds = answerTime
}

func (answerData *QuestionData) Start() {
	now := time.Now()
	answerData.Started = &now
}

func (answerData *QuestionData) Answer(userId id.ID, answer models.AnswerMessage) error {
	if answerData.AnsweredBy != nil {
		return ErrAlreadyAnswered
	}
	answerData.AnsweredBy = &userId
	now := time.Now()
	answerData.Answered = &now
	isCorrect := answerData.Question.Answer.AnswerData.IsCorrect(answer)
	answerData.AnsweredCorrectly = &isCorrect
	return nil
}

func (answerData *QuestionData) Reset() {
	answerData.AnsweredBy = nil
	answerData.Answered = nil
	answerData.Started = nil
}
