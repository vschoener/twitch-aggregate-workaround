package core

import "time"

// SubscriptionSummary is a structure containing subscription information
type SubscriptionSummary struct {
	ID          int64
	CreatedAt   time.Time
	SubPlan     string
	SubPlanName string
	Users       []UserSubscription
}

// UserSubscription contains user subscription information
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
