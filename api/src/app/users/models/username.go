package models

import (
	"database/sql/driver"
	"errors"
)

type UserName string

func (userName UserName) Valid() error {
	if string(userName) == "" {
		return errors.New("user_name cannot be empty")
	}
	return nil
}

func (UserName) GormDataType() string { return "varchar(48)" }
func (vo UserName) GetValue() (driver.Value, error) {
	if err := vo.Valid(); err != nil {
		return nil, err
	}
	return string(vo), nil
}
