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

func (r *Repository) applyFilter(db *gorm.DB, filter storage.QueryFilter) *gorm.DB {
	for _, dateRange := range filter.Ranges {
		if nil != filter.DateStart {
			db = db.Where(dateRange.DateField+" >= ?", dateRange.DateStart.Format(storage.SIMPLEFORMATSQL))
		}
		if nil != filter.DateEnd {
			db = db.Where(dateRange.DateField+" <= DATE_ADD(?, INTERVAL 1 DAY)", dateRange.DateEnd.Format(storage.SIMPLEFORMATSQL))
		}
	}

	if nil != filter.Exclude && len(filter.Exclude) > 0 {
		for field, values := range filter.Exclude {
			db = db.Where(field+" NOT IN (?)", values)
		}
	}
	if nil != filter.Include && len(filter.Include) > 0 {
		for field, values := range filter.Include {
			db = db.Where(field+" IN (?)", values)
		}
	}

	return db
}
