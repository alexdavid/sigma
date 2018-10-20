package sigma

import "time"

// Cocoa epoch is Jan 1st 2001 not 1970 so add number of seconds in 31 years
const cocoa_unix_epoc_diff int64 = 978285600

func cocoaTimestampToTime(timestamp int) time.Time {
	return time.Unix(int64(timestamp)+cocoa_unix_epoc_diff, 0)
}

func timeToCocoaTimestamp(time time.Time) int64 {
	return time.Unix() - cocoa_unix_epoc_diff
}
