package utils

import "math"

func isSameDay(t1Stamp int64, t2Stamp int64, restTime int64) bool {
	dayPerSeconds := int64(24 * 3600)
	diff1 := (t1Stamp - restTime*3600) / dayPerSeconds
	diff2 := (t2Stamp - restTime*3600) / dayPerSeconds
	if math.Abs(float64(diff1-diff2)) > 0 {
		return false
	}
	return true
}
