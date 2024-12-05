package repo

import (
	"lib/app/ws/models"
	"lib/common/id"
	"lib/services"
	"log"
)

type metadataRepo struct{}

func (metadataRepo) GetBySocket(socketId id.ID) (models.Metadata, error) {
	db, freeDB := services.Db()
	defer freeDB()
	var model models.Metadata
	err := db.Where("socket_id = ?", socketId).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (metadataRepo) GetByUserId(userId id.ID) (models.Metadata, error) {
	db, freeDB := services.Db()
	defer freeDB()
	var model models.Metadata
	err := db.Where("session->>'user_id' = ?", userId).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (metadataRepo) Add(model models.Metadata) error {
	db, freeDB := services.Db()
	defer freeDB()
	err := db.Create(model).Error
	return err
}

func (metadataRepo) Delete(id id.ID) {
	db, freeDB := services.Db()
	defer freeDB()
	tx := db.Where("socket_id = ?", id).Unscoped().Delete(&models.Metadata{})
	if tx.RowsAffected == 0 {
		log.Printf("no rows where removed from models metadata where id = %v", id)
	}
}

func (metadataRepo) DeleteAll() {
	db, freeDb := services.Db()
	defer freeDb()
	db.Where("socket_id = socket_id").Unscoped().Delete(&models.Metadata{})
}
