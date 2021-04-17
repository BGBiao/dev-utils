package time

import (
	"time"
)

const timeFormatCommon = "2006-01-02T15:04:05"

func UnixTsToString(ts int64) string {
	return time.Unix(ts, 0).Format(timeFormatCommon)
}

func StringTsToUnix(date string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")

	t, paserTimeErr := time.ParseInLocation(timeFormatCommon, date, loc)

	if paserTimeErr != nil {
		return 0
	}

	return t.Unix()

}

func NowTime() string {
	return time.Now().Format(timeFormatCommon)
}

func NowAfterTime(second int64) string {
	return time.Now().Add(time.Duration(second) * time.Second).Format(timeFormatCommon)
}
