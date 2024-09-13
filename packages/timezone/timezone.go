package timezone

import "time"

// LocalTimeZone is the timezone of the application.
const LocalTimeZone = "Asia/Bangkok"

func NewAsiaBangkok() *time.Location {
	timeLocal, err := time.LoadLocation(LocalTimeZone)
	if err != nil {
		panic(err)
	}

	return timeLocal
}
