package model

import (
	"time"

	"github.com/wonderstream/twitch/core/model"
)

const (
	// UserTable database table
	UserTable = "user"
)

// User is the storage manager
type User struct {
	ID      int64
	DateAdd time.Time
	model.User
}
