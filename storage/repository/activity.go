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
}

// CCVInformation contains CCV information as AVG and Max
type CCVInformation struct {
	MaxCCV   int64
	AvgCCV   int64
	TotalCCV int64
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
// TODO: Add PRIVMSG in the check if the user is already on the channel is talking                     ^
func (r ActivityStorageRepository) GetWatchedTime(name string, filters storage.QueryFilter) int64 {
	filters.Include = map[string][]string{"type": {"JOIN", "PART"}}

	log.Println(filters)
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

// timeActivity as "YYYY-MM-DD HH:ii:ss" for example
type timeActivity string

// username is a string
type username string

type usersActivity map[username]*model.Activity

// usersActivities contains a list of timeActivity which contains username list
type timeActivities map[timeActivity]usersActivity

func getUsersStillConnected(connectedList usersActivity, disconnectedList usersActivity) usersActivity {
	usersStillConnected := make(usersActivity)
	// First copy the connected list to the new one
	for username, activity := range connectedList {
		usersStillConnected[username] = activity
	}

	// Then delete entry for the disconnected user
	for username := range disconnectedList {
		delete(usersStillConnected, username)
	}
	return usersStillConnected
}

// GetCCVInformation compute concurrent user information as Max and AVG
func (r ActivityStorageRepository) GetCCVInformation(name string, filters storage.QueryFilter) CCVInformation {
	ccvInfo := CCVInformation{}
	activities := r.GetChannelActivities(name, filters)

	// connectedUsers contains a list of user be disconnected at a time
	timesConnectedUsers := timeActivities{}
	// disconnectedUsers contains a list of user be disconnected at a time
	timesDisconnectedUsers := timeActivities{}
	lastMinuteTime := time.Time{}
	oldDates := []timeActivity{}

	for _, activity := range activities {
		dateByMinute := timeActivity(activity.DateTime.Format(storage.SQLTimeFormatByMinute))
		// Reverse to use time logic (as we can't set specific time to specific value)
		dateByMinuteTime, _ := time.Parse(storage.SQLTimeFormatByMinute, string(dateByMinute))

		if dateByMinuteTime.Sub(lastMinuteTime).Minutes() >= 1 {
			// Allocate memory map for the new time
			timesConnectedUsers[dateByMinute] = make(usersActivity)
			timesDisconnectedUsers[dateByMinute] = make(usersActivity)

			if !lastMinuteTime.IsZero() {
				// lastMinuteTime contains the old / last time by minute
				lastDateByMinute := timeActivity(lastMinuteTime.Format(storage.SQLTimeFormatByMinute))
				oldDates = append(oldDates, lastDateByMinute)

				// Build a new list of connected user
				timesConnectedUsers[dateByMinute] = getUsersStillConnected(
					timesConnectedUsers[lastDateByMinute],
					timesDisconnectedUsers[lastDateByMinute],
				)
			}
			// Store anyway the last date by minute
			lastMinuteTime = dateByMinuteTime
		}

		currentUser := username(activity.Username)
		_, userFound := timesConnectedUsers[dateByMinute][currentUser]

		// If user is not found whatever the type is, he should be added to the
		// list of connected users
		if !userFound {
			timesConnectedUsers[dateByMinute][currentUser] = &activity
		}

		// In case it's a PART, we have more to do
		// - ADD use to the disconnected users connectedList
		// - If not found, check back the old lists and review CCV
		if activity.Type == "PART" {
			timesDisconnectedUsers[dateByMinute][currentUser] = &activity
			if !userFound {
				// Review old list and CCV until we found back from
				// recent date to the oldest one
				for i := len(oldDates) - 1; i >= 0; i-- {
					oldDateByMinuteKey := oldDates[i]
					// Check if the user is not into an old list
					// ex JOIN -> PART JOIN ... BOT DISCO | network issue ... PART
					if _, found := timesConnectedUsers[oldDateByMinuteKey][currentUser]; found {
						// The user exists in an old list, stop the CCV recount
						break
					}
					// User is still missing, let's review the CCV
					timesConnectedUsers[oldDateByMinuteKey][currentUser] = nil
				}
			}
		}
	}

	// Let's fill our CCV info
	for _, users := range timesConnectedUsers {
		totalCCV := len(users)
		ccvInfo.TotalCCV += int64(totalCCV)
		if ccvInfo.MaxCCV < int64(totalCCV) {
			ccvInfo.MaxCCV = int64(totalCCV)
		}
	}

	if totalDate := int64(len(timesConnectedUsers)); totalDate > 0 {
		ccvInfo.AvgCCV = int64(ccvInfo.TotalCCV / totalDate)
	}

	return ccvInfo
}
