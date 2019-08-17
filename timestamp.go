package sigma

import "time"

// Cocoa epoch is Jan 1st 2001 not 1970 so add number of seconds in 31 years
const cocoaUnixEpocDiff int64 = 978285600
const nanosecondsInSecond int64 = 1000000000

func cocoaTimestampToTime(timestamp int64) time.Time {
	if timestamp > 1000000000000 {
		// If timestamp is bigger than 1000000000000 we can safely assume it's in nanoseconds
		// Older versions of macos use seconds, newer use nanoseconds
		timestamp = timestamp / nanosecondsInSecond
	}

	return time.Unix(timestamp+cocoaUnixEpocDiff, 0)
}
