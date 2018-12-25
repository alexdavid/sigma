package sigma

import "time"

// Cocoa epoch is Jan 1st 2001 not 1970 so add number of seconds in 31 years
const cocoa_unix_epoc_diff int64 = 978285600
const nanoseconds_in_second int64 = 1000000000

func cocoaTimestampToTime(timestamp int64) time.Time {
	if timestamp > 1000000000000 {
		// If timestamp is bigger than 1000000000000 we can safely assume it's in nanoseconds
		// Older versions of macos use seconds, newer use nanoseconds
		timestamp = timestamp / nanoseconds_in_second
	}

	return time.Unix(timestamp+cocoa_unix_epoc_diff, 0)
}
