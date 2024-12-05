package models

import "errors"

var ErrAnswerTimeOutOfScope error = errors.New("answer time must be between 10 and 120 seconds")

type AnswerTimeInSeconds int

var DefaultAnswerTime AnswerTimeInSeconds = 30

func (time AnswerTimeInSeconds) Valid() error {
	if int(time) < 10 || int(time) > 120 {
		return ErrAnswerTimeOutOfScope
	}
	return nil
}
