package repository

import (
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
)

// PrecomputeChannelRepository handles channel database query
type PrecomputeChannelRepository struct {
	*Repository
}

// NewPrecomputeChannelRepository return a credential repository
func NewPrecomputeChannelRepository(db *storage.Database, l logger.Logger) PrecomputeChannelRepository {
	commonRepository := NewRepository(db, l)
	r := PrecomputeChannelRepository{
		Repository: commonRepository,
	}

	return r
}

// Add new entry to the database (Shortcut ways)
func (p PrecomputeChannelRepository) Add(entity model.PrecomputedChannel) error {
	err := p.Database.Gorm.Create(&entity).Error

	return err
}
