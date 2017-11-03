package repository

import (
	"github.com/jinzhu/gorm"
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

// CheckErr check any error or log
func (r *Repository) CheckErr(err error) bool {
	if nil != err {
		r.Database.Logger.LogInterface(err)

		return false
	}

	return true
}

func (r *Repository) applyFilter(db *gorm.DB, filter storage.QueryFilter) {
	if nil != filter.DateStart {
		db = db.Where(filter.DateField+" >= ?", filter.DateStart.Format(storage.SIMPLEFORMATSQL))
	}
	if nil != filter.DateEnd {
		db = db.Where(filter.DateField+" <= DATE_ADD(?, INTERVAL 1 DAY)", filter.DateEnd.Format(storage.SIMPLEFORMATSQL))
	}
}
