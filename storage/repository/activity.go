package repository

import (
	"log"
	"time"

	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
)

// ActivityStorageRepository handles channel database query
type ActivityStorageRepository struct {
	*Repository
	cache         bool
	CacheActivity CacheActivity
}

// CacheActivity used to manage cache for any query from this repository
type CacheActivity struct {
	Activities []model.Activity
}

// SetCacheSet is used to cache redundant data to avoid extra
// process consumption
func (r ActivityStorageRepository) SetCacheSet(state bool) {
	r.cache = state
}

// NewActivityStorageRepository return a credential repository
func NewActivityStorageRepository(db *storage.Database, l logger.Logger) ActivityStorageRepository {
	commonRepository := NewRepository(db, l)
	r := ActivityStorageRepository{
		Repository: commonRepository,
	}

	return r
}

// GetChannelActivities list of channel
func (r ActivityStorageRepository) GetChannelActivities(name string, filters storage.QueryFilter) []model.Activity {
	activities := []model.Activity{}
	db := r.Database.Gorm.
		Model(&model.Activity{}).
		Where(`channel = ?`,
			name,
		).Order("datetime ASC")

	filters.DateField = "datetime"
	db = r.applyFilter(db, filters)
	db.Find(&activities)

	return activities
}

// GetWatchedTime list of channel
func (r ActivityStorageRepository) GetWatchedTime(name string, filters storage.QueryFilter) int64 {
	activities := r.GetChannelActivities(name, filters)
	var seconds int64
	joinUsersActivity := map[string]model.Activity{}

	var olderDate time.Time
	for _, activity := range activities {
		// Save the older date
		if olderDate.IsZero() {
			olderDate = activity.DateTime
		}

		if activity.Type == "JOIN" {
			// /!\ Maybe the user could exist and JOIN again because bot could be
			//     disconnected ?
			joinUsersActivity[activity.Username] = activity
		} else if activity.Type == "PART" {
			var dateToCompare time.Time
			if _, found := joinUsersActivity[activity.Username]; found {
				// If the user is in our list we're good, but it some case
				// the user could not be in the list (bot not connected for example)
				dateToCompare = joinUsersActivity[activity.Username].DateTime
				delete(joinUsersActivity, activity.Username)
			} else {
				// Then we take the older date of the range
				dateToCompare = olderDate
			}
			diff := activity.DateTime.Sub(dateToCompare)
			seconds += int64(diff.Seconds())
		}
	}

	return seconds
}

// GetUniqueViewers return a number of unique viewers from a channel
func (r ActivityStorageRepository) GetUniqueViewers(name string, filters storage.QueryFilter) int64 {
	type Result struct {
		TotalUniqueViewers int64
	}

	result := Result{}
	db := r.Database.Gorm.
		Model(&model.Activity{}).
		Select("COUNT(DISTINCT username) total_unique_viewers").
		Where(`channel = ?`,
			name,
		).Order("username")

	filters.DateField = "datetime"
	db = r.applyFilter(db, filters)

	err := db.Scan(&result).Error

	if nil != err {
		log.Println(err)
	}

	return result.TotalUniqueViewers
}
