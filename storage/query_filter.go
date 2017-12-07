package storage

import "time"

const (
	// SIMPLEFORMATSQL to format time.Time
	SIMPLEFORMATSQL = "2006-01-02"

	// SQLTimeFormatByMinute use as layout
	SQLTimeFormatByMinute = SIMPLEFORMATSQL + " 15:04:00"
)

// DateFilter provides fields to filter date
type DateFilter struct {
	DateField string
	DateStart time.Time
	DateEnd   time.Time
}

// UseDateField allows to change the date field used on the fly
func (d DateFilter) UseDateField(name string) DateFilter {
	d.DateField = name

	return d
}

// QueryFilter used to filter repository query
type QueryFilter struct {
	Ranges []DateFilter
	DateFilter
	Exclude map[string][]string
	Include map[string][]string
}
