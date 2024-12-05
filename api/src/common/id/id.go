package id

import (
	"database/sql/driver"
	"errors"

	"github.com/google/uuid"
)

type ID string

func (id ID) Valid() error {
	if _, err := uuid.Parse(string(id)); err != nil {
		return errors.New("id is not valid")
	}
	return nil
}

func (id ID) GetValue() (driver.Value, error) {
	if err := id.Valid(); err != nil {
		return nil, err
	}
	return string(id), nil
}
func (ID) GormDataType() string { return "varchar(36)" }

func New() ID {
	id := ID(uuid.NewString())
	return id
}
