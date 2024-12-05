package models

type AnswerMessage string

func (AnswerMessage) Valid() error {
	return nil
}
