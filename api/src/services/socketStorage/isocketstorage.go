package socketstorage

import (
	"lib/common/id"
)

type SocketStorage interface {
	Get(id.ID) Ws
	Add(Ws) id.ID
	OnConnect(func(id id.ID, ws Ws))
}
