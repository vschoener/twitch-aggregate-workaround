package repository

import (
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
)

// Repository is a (parent) struct containing common requirement
type Repository struct {
	Database *storage.Database
	Logger   logger.Logger
}

// NewRepository constructor
func NewRepository(d *storage.Database, l logger.Logger) *Repository {
	return &Repository{
		Database: d,
		Logger:   l,
	}
}
