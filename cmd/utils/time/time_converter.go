package time

import "time"

func TimeRFC3339From(date string) *time.Time {
	d, _ := time.Parse(time.RFC3339, date)
	return &d
}
