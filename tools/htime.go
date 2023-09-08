package tools

import "time"

var (
	local, _  = time.LoadLocation("Asia/Chongqing") //服务器设置的时区
	TIMELOCAL = local
)

// NowDate 当前时区标准日期
func NowDate() string {
	return time.Now().In(TIMELOCAL).Format("2006-01-02 15:04:05")
}

// YesterdayDate 当前昨天日期(Y-m-d)
func YesterdayDate() string {
	yesterday := time.Now().In(TIMELOCAL).AddDate(0, 0, -1)
	return yesterday.Format("2006-01-02")
}

// StrToTimeStampLocal 字符串转为unix时间戳
func StrToTimeStampLocal(timeStr string) int64 {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, local)
	return t.Unix()
}

// StrToTimeStamp 字符串转为unix时间戳
func StrToTimeStamp(timeStr string) int64 {
	t, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	return t.Unix()
}
