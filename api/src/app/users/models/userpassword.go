package models

import (
	"database/sql/driver"
	"errors"
)

type UserPassword string

func (password UserPassword) Valid() error {
	if string(password) == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}

func (vo UserPassword) GetValue() (driver.Value, error) {
	if err := vo.Valid(); err != nil {
		return nil, err
	}
	return string(vo), nil
}
