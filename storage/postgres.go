package storage

import (
	"fmt"
	"manifest-craft/database"
	"manifest-craft/models"
)

type PostgressStorage struct{}

func NewPostgressStorage() *PostgressStorage {
	return &PostgressStorage{}
}

func (s *PostgressStorage) Get(name string) *models.Manifest {
	manifest := models.Manifest{}

	fmt.Println("Getting manifest from database...")

	if err := database.DB.Db.Preload("Components").Where("name = ?", name).First(&manifest).Error; err != nil {
		fmt.Println("Error getting manifest from database:", err)
		return nil
	}

	return &manifest
}
