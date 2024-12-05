package services

import (
	"lib/common/servicepool"

	"gorm.io/gorm"
)

var _db servicepool.ServicePool[*gorm.DB]

func SetDb(db servicepool.ServicePool[*gorm.DB]) {
	_db = db
}

func Db() (*gorm.DB, func()) {
	if _db == nil {
		panic("Cannot access Db service. Db is not defined")
	}
	return _db.Get()
}
