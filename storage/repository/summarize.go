package repository

import (
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
)

// SummarizeRepository handles channel database query
type SummarizeRepository struct {
	*Repository
}

// NewSummarizeRepository return a credential repository
func NewSummarizeRepository(db *storage.Database, l logger.Logger) SummarizeRepository {
	commonRepository := NewRepository(db, l)
	r := SummarizeRepository{
		Repository: commonRepository,
	}

	return r
}
