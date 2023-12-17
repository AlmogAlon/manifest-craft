package storage

import "manifest-craft/models"

type Storage interface {
	Get(name string) *models.Manifest
	GetComponent(name string) *models.Component
}
