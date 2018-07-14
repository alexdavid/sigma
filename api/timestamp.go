package api

import "time"

// Cocoa epoch is Jan 1st 2001 not 1970 so add number of seconds in 31 years
const COCOA_UNIX_EPOC_DIFF int64 = 978285600

func cocoaTimestampToTime(timestamp int) time.Time {
	return time.Unix(int64(timestamp)+COCOA_UNIX_EPOC_DIFF, 0)
}

func timeToCocoaTimestamp(time time.Time) int64 {
	return time.Unix() - COCOA_UNIX_EPOC_DIFF
}
