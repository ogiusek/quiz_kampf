package models

import (
	"errors"
)

type QuestionText string

func (text QuestionText) Valid() error {
	if len(string(text)) == 0 {
		return errors.New("`question` cannot be empty")
	}
	return nil
}
