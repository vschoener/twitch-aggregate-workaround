package model

import "time"

// Subscription the model from Twitch API
type Subscription struct {
	ID          int64
	CreatedAt   time.Time
	SubPlan     string
	SubPlanName string
	Users       []UserSubscription
}

// UserSubscription is the model from Twitch API
type UserSubscription struct {
	ID          int64
	Bio         string
	CreatedAt   time.Time
	DisplayName string
	Logo        string
	Name        string
	Type        string
	UpdatedAt   time.Time
}
