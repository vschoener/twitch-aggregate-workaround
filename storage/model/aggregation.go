package model

import "time"

const (
	aggregationTable = "aggregation"
)

// Aggregation contains last request on different ressource
type Aggregation struct {
	ChannelID               int64
	LastStreamSession       time.Time
	LastChannelSummary      time.Time
	LastSubscriptionSummary time.Time
	LastStream              time.Time
	LastUsers               time.Time
}
