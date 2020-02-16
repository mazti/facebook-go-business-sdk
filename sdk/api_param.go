package sdk

import (
	"time"
)

const (
	TimeRangeKey = "time_range"
)

type TimeRange struct {
	Since string `json:"since"`
	Until string `json:"until"`
}

func CreateTimeRange(since time.Time, until time.Time) TimeRange {
	return TimeRange{
		Since: since.Format(TimeFormat),
		Until: until.Format(TimeFormat),
	}
}
