package models

import (
	"errors"
	"strings"
)

type AnswerTextType struct {
	CorrectAnswers []AnswerMessage `json:"correct_answers"`
	CaseSensitive  bool            `json:"case_sensitive"`
}

func (answer AnswerTextType) IsCorrect(message AnswerMessage) bool {
	for _, correctMessage := range answer.CorrectAnswers {
		if (answer.CaseSensitive && message == correctMessage) ||
			(!answer.CaseSensitive && strings.EqualFold(string(message), string(correctMessage))) {
			return true
		}
	}
	return false
}

func NewAnswerText(correctAnswers []AnswerMessage, caseSesitive bool) Answer {
	return Answer{
		AnswerType: AnswerText,
		AnswerData: AnswerTextType{
			CorrectAnswers: correctAnswers,
			CaseSensitive:  caseSesitive,
		},
	}
}

func (options AnswerTextType) Valid() error {
	if len(options.CorrectAnswers) == 0 {
		return errors.New("options cannot be empty")
	}
	for _, answer := range options.CorrectAnswers {
		if err := answer.Valid(); err != nil {
			return err
		}
	}
	return nil
}
