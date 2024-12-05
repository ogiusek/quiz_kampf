package socketstorage

import (
	"lib/common/id"
	"sync"
)

type socketStorage struct {
	socketsMutex   sync.Mutex
	sockets        map[id.ID]Ws
	onConnectMutex sync.Mutex
	onConnect      []func(id.ID, Ws)
}

func (storage *socketStorage) Get(id id.ID) Ws {
	ws, found := storage.sockets[id]
	if !found {
		return nil
	}
	return ws
}

func (storage *socketStorage) OnConnect(fn func(id.ID, Ws)) {
	storage.socketsMutex.Lock()
	defer storage.socketsMutex.Unlock()
	storage.onConnect = append(storage.onConnect, fn)
}

func (storage *socketStorage) Add(ws Ws) id.ID {
	storage.onConnectMutex.Lock()
	defer storage.onConnectMutex.Unlock()
	id := id.New()
	storage.sockets[id] = ws
	for _, fn := range storage.onConnect {
		fn(id, ws)
	}
	return id
}

func NewStorage() SocketStorage {
	storage := &socketStorage{
		sockets: map[id.ID]Ws{},
	}

	return storage
}
