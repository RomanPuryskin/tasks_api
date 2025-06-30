package storage

import (
	"github.com/tasks_api/pkg/models"
)

var DB *models.Storage

func NewStorage() (*models.Storage, error) {
	DB = &models.Storage{
		Db: make(map[string]*models.Task),
	}

	return DB, nil
}
