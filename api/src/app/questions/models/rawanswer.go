package models

import (
	"errors"
	"lib/common/mapping"
	"lib/common/valid"
)

type RawAnswer struct {
	AnswerType AnswerType `json:"answer_type"`
	AnswerData any        `json:"answer_data"`
}

func (rawAnswer RawAnswer) ToAnswer() (Answer, error) {
	answer := Answer{
		AnswerType: rawAnswer.AnswerType,
	}

	var finalValue AnswerData
	var err error

	switch answer.AnswerType {
	case AnswerOptions:
		var value AnswerOptionsType
		err = mapping.MapToStruct(rawAnswer.AnswerData, &value)
		finalValue = value
	case AnswerText:
		var value AnswerTextType
		err = mapping.MapToStruct(rawAnswer.AnswerData, &value)
		finalValue = value
	default:
		err = errors.New("invalid answer type")
	}

	if err != nil {
		return Answer{}, err
	}
	answer.AnswerData = finalValue

	return answer, nil
}

func (rawAnswer RawAnswer) Valid() error {
	answer, err := rawAnswer.ToAnswer()
	if err != nil {
		return err
	}

	return valid.Valid(answer)
}
