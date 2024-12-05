package models

import "errors"

type AnswerOptionsType struct {
	Answers []AnswerMessage `json:"answers"`
	Correct int             `json:"correct"`
}

func (answer AnswerOptionsType) IsCorrect(message AnswerMessage) bool {
	for i, mess := range answer.Answers {
		if mess == message {
			if i == answer.Correct {
				return true
			}
			break
		}
	}
	return false
}

func NewAnswerOptions(Answers []AnswerMessage) Answer {
	return Answer{
		AnswerType: AnswerOptions,
		AnswerData: AnswerOptionsType{
			Answers: Answers,
		},
	}
}

func (options AnswerOptionsType) Valid() error {
	if len(options.Answers) < 2 {
		return errors.New("answers have to be at least 2")
	}
	for _, answer := range options.Answers {
		if err := answer.Valid(); err != nil {
			return err
		}
	}
	if options.Correct < 0 || options.Correct > len(options.Answers) {
		return errors.New("correct answer has to be between 0 and answers amount")
	}
	return nil
}
