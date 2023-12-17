package storage

import (
	"manifest-craft/config"
	"manifest-craft/models"

	log "github.com/sirupsen/logrus"
)

type PostgressStorage struct{}

func NewPostgressStorage() *PostgressStorage {
	return &PostgressStorage{}
}

func (s *PostgressStorage) Get(name string) *models.Manifest {
	manifest := models.Manifest{}

	log.Info("Getting manifest from database...")

	if err := config.DB.Db.Preload("Components").Where("name = ?", name).First(&manifest).Error; err != nil {
		log.Info("Error getting manifest from database:", err)
		return nil
	}

	return &manifest
}
