package utils

import "time"

// MillTimeStampToTime convert ms timestamp to time.Time
func MillTimeStampToTime(timestamp int64) time.Time {
	second := timestamp / 1000
	nano := timestamp % 1000 * 1000000
	return time.Unix(second, nano)
}

// SecondTimeStampToTime convert s timestamp to time.Time
func SecondTimeStampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
