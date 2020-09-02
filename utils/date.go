package utils

//IsSameDay is used to judge 2 timestamp is sameday
//curTime : current time
//LastTime : Want compare time
//restTime:the hour reset a day
//zone : the time zone ex: UTC+0 zone =0
func IsSameDay(curTime int64, lastTime int64, restTime int64, zone int64) bool {
	return betweenDays(curTime-3600*restTime, lastTime-3600*restTime, zone) == 0
}

func betweenDays(t1 int64, t2 int64, zone int64) int64 {
	secPerDay := int64(3600 * 24)
	return (t1+zone*3600)/secPerDay - (t2+zone*3600)/secPerDay
}
