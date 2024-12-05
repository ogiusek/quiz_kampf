package repo

import (
	"lib/app/ws/models"
	"lib/common/id"
)

type MetadataGetRepo interface {
	GetBySocket(socketId id.ID) (models.Metadata, error)
	GetByUserId(userId id.ID) (models.Metadata, error)
}

type MetadataModifyRepo interface {
	Add(models.Metadata) error
	Delete(id id.ID)
	DeleteAll()
}

type MetadataRepo interface {
	MetadataGetRepo
	MetadataModifyRepo
}

func GetMetadataGetRepo() MetadataGetRepo       { return metadataRepo{} }
func GetMetadataModifyRepo() MetadataModifyRepo { return metadataRepo{} }
func GetMetadataRepo() MetadataRepo             { return metadataRepo{} }
