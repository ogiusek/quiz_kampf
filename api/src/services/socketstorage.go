package services

import (
	socketstorage "lib/services/socketStorage"
)

// var _socketStorage servicepool.ServicePool[socketstorage.SocketStorage] =

//	func SetSocketStorage(socketStorage servicepool.ServicePool[socketstorage.SocketStorage]) {
//		_socketStorage = socketStorage
//	}

var _socketStorage socketstorage.SocketStorage = socketstorage.NewStorage()

func SocketStoage() (socketstorage.SocketStorage, func()) {
	// if _socketStorage == nil {
	// 	panic("Cannot access socket storage. socket storage is not defined")
	// }
	// return _socketStorage.Get()
	return _socketStorage, func() {}
}
