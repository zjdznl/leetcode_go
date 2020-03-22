package leetcode

import (
	"math"
	"time"
)

func ParseTimeStrToTime(timeS string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timeS, time.Local)
	if err != nil {
		return time.Now()
	}
	return t
}

func daysBetweenDates(date1 string, date2 string) int {
	date1 += " 00:00:00"
	date2 += " 00:00:00"
	t1 := ParseTimeStrToTime(date1)
	t2 := ParseTimeStrToTime(date2)
	return int(math.Ceil(math.Abs(t1.Sub(t2).Hours()) / 24))
}
