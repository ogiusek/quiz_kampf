package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"lib/common/valid"
)

type AnswerData interface {
	IsCorrect(AnswerMessage) bool
}

type Answer struct {
	AnswerType AnswerType `json:"answer_type"`
	AnswerData AnswerData `json:"answer_data"`
}

func (Answer) GormDataType() string { return "jsonb" }
func (vo Answer) GetValue() (driver.Value, error) {
	if err := valid.Valid(vo); err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(vo)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (answer *Answer) Scan(value interface{}) error {
	valueByte, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan value")
	}
	var rawAnswer RawAnswer
	err := json.Unmarshal(valueByte, &rawAnswer)
	if err != nil {
		return err
	}
	result, err := rawAnswer.ToAnswer()
	if err != nil {
		return err
	}
	*answer = result
	return nil
}
