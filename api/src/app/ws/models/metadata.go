package models

import (
	"lib/app/users/dto"
	"lib/common/id"
)

type Metadata struct {
	SocketId id.ID       `gorm:"column:socket_id;not null"`
	Session  dto.Session `gorm:"column:session;not null"`
}
