package models

import (
	"database/sql/driver"
	"lib/services/filestorage"
)

type UserImage filestorage.File

func (e UserImage) GormDataType() string { return filestorage.File(e).GormDataType() }
func (vo UserImage) GetValue() (driver.Value, error) {
	return filestorage.File(vo).GetValue()
}
