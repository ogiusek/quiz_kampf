package models

import (
	"database/sql/driver"
	"errors"
)

type AnswerType int

const (
	AnswerOptions AnswerType = iota + 1
	AnswerText
	// AnswerLocation
	// AnswerColor
)

func (vo AnswerType) Valid() error {
	if int(vo) < 1 || int(vo) > 2 {
		return errors.New("answer type must be between 1 and 2")
	}
	return nil
}

func (AnswerType) GormDataType() string { return "tinyint" }
func (vo AnswerType) GetValue() (driver.Value, error) {
	if err := vo.Valid(); err != nil {
		return nil, err
	}
	return int(vo), nil
}
