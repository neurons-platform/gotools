package utils

import (
	"github.com/araddon/dateparse"
	"time"
)

func GetMilliTimeStamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func MilliTimeStamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetDatetime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05 PM")
}

func Sleep(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func ParserStringToTime(str string) time.Time {
	t, err := dateparse.ParseLocal(str)
	// t, err := dateparse.ParseFormat(str)
	Throw(err)
	return t
}

func GetLastNMinute(n int64) int64 {
	t1 := MilliTimeStamp()
	t2 := t1 % 60000
	return t1 - t2 - 60000*n
}
