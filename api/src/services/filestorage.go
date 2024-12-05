package services

import (
	"lib/common/servicepool"
	"lib/services/filestorage"
)

var _fileStorage servicepool.ServicePool[filestorage.FileStorage]

func SetFileStorage(fileStorage servicepool.ServicePool[filestorage.FileStorage]) {
	_fileStorage = fileStorage
}

func FileStoage() (filestorage.FileStorage, func()) {
	if _fileStorage == nil {
		panic("Cannot access file storage. file storage is not defined")
	}
	return _fileStorage.Get()
}
