package tools

import "time"

const (
	timeKey = "2006-01-02 15:04:05"
	dateKey = "20060102"
)

func UnixToDate(timestamp int) string {
	return time.Unix(int64(timestamp), 0).Format(timeKey)
}

func DateToUnix(date string) int64 {
	t, err := time.ParseInLocation(timeKey, date, time.Local)
	if err != nil {
		return -1
	}
	return t.Unix()
}

func GetDate() string {
	return time.Now().Format(dateKey)
}
