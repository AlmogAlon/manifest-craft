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

func (s *PostgressStorage) GetComponent(source string) *models.Component {
	component := models.Component{}

	log.Info("Getting components from database...")

	if err := config.DB.Db.Where("source = ?", source).First(&component).Error; err != nil {
		log.Info("Error getting components from database:", err)
		return nil
	}

	return &component
}
