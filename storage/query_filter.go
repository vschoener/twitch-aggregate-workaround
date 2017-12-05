package storage

import "time"

const (
	// SIMPLEFORMATSQL to format time.Time
	SIMPLEFORMATSQL = "2006-01-02"
)

// QueryFilter used to filter repository query
type QueryFilter struct {
	Ranges []struct {
		DateField string
		DateStart *time.Time
		DateEnd   *time.Time
	}
	DateStart *time.Time
	DateEnd   *time.Time
	DateField string
	Exclude   map[string][]string
	Include   map[string][]string
}
