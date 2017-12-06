package repository

import (
	"time"

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

func (r *Repository) applyDateFilter(db *gorm.DB, field string, start time.Time, end time.Time) *gorm.DB {
	if !start.IsZero() {
		db = db.Where(field+" >= ?", start.Format(storage.SIMPLEFORMATSQL))
	}

	if !end.IsZero() {
		db = db.Where(field+" <= DATE_ADD(?, INTERVAL 1 DAY)", end.Format(storage.SIMPLEFORMATSQL))
	}

	return db
}

func (r *Repository) applyFilter(db *gorm.DB, filter storage.QueryFilter) *gorm.DB {
	if len(filter.Ranges) > 0 {
		for _, dateRange := range filter.Ranges {
			db = r.applyDateFilter(db, dateRange.DateField, dateRange.DateStart, dateRange.DateEnd)
		}
	} else {
		db = r.applyDateFilter(db, filter.DateField, filter.DateStart, filter.DateEnd)
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
