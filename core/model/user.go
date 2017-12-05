package model

import (
	"time"

	"github.com/wonderstream/twitch/core"
)

// Users manager
type Users struct {
	*core.Request
}

// User information
type User struct {
	DisplayName string    `json:"display_name"`
	ID          int64     `json:"_id,string"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Bio         string    `json:"bio"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Logo        string    `json:"logo"`
}
