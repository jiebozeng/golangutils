// Based on Carcon development
// 2006-01-02 15:04:05
package timeutils

import (
	"github.com/golang-module/carbon"
	"strconv"
	"strings"
	"time"
)

/*
*
返回当前日期
格式 20060102
*/
func GetNowDateInt20060815() int32 {
	now := time.Now()
	return int32(now.Year()*10000 + int(now.Month())*100 + now.Day())
}

/*
*
返回当前时间
格式 150405
15点04分05秒
*/
func GetNowTimeInt150405() int32 {
	now := time.Now()
	return int32(now.Hour()*10000 + int(now.Minute())*100 + now.Second())
}

// 获取当前时间 返回的格式 2006-01-02 15:04:05
func GetNowTime() string {
	return carbon.Now().ToDateTimeString()
}

// 时间差 当前时间 - 参数时间 相差多少秒
// 参数的格式 2020-08-04 14:14:15
func HowLongFromNow(value string) int64 {
	return carbon.Parse(value).DiffInSeconds(carbon.Parse(GetNowTime()))
}

// 获取当前时间 15:04:05 小时:分:秒
func GetNowTimeString() string {
	return carbon.Now().ToTimeString()
}

// 返回当前时间的时间戮 秒
func GetNowTimestampInt64() int64 {
	return carbon.Now().Timestamp()
}

// 参数的格式 2006-01-02 15:04:05 返回 true or false
func IsYesterday(value string) bool {
	return carbon.Parse(value).IsYesterday()
}

// 参数15:04:05 HH:MM:SS
// 返回秒数
func HourAndMinuteAndSecondToSecond(t string) int64 {
	tt := strings.Split(t, ":")
	if len(tt) < 3 {
		return 0
	}
	h, _ := strconv.ParseInt(tt[0], 10, 64)
	m, _ := strconv.ParseInt(tt[1], 10, 64)
	s, _ := strconv.ParseInt(tt[2], 10, 64)
	return h*3600 + m*60 + s
}

// 15小时45分钟32秒 后的时间
// 参数 beginTime 2020-08-04 14:14:15
// 返回格式 时间戳 秒
func HourAndMinuteAndSecondLater(beginTime string, hms string) int64 {
	tt := strings.Split(hms, ":")
	if len(tt) < 3 {
		return 0
	}
	return carbon.Parse(beginTime).AddDuration(tt[0] + "h" + tt[1] + "m" + tt[2] + "s").Timestamp()
}

// 5小时45分钟32秒 后的时间
// 参数 beginTime 2006-01-02 15:04:05
// 返回格式 2006-01-02 20:49:37
func HourAndMinuteAndSecondLaterString(beginTime string, hms string) string {
	tt := strings.Split(hms, ":")
	if len(tt) < 3 {
		return ""
	}
	return carbon.Parse(beginTime).AddDuration(tt[0] + "h" + tt[1] + "m" + tt[2] + "s").ToDateTimeString()
}

// value second 秒数
// return second changeto 15:45:32 15小时45分钟32秒
func SecondToHourAndMinuteAndSecond(value int64) string {
	var hour int64 = value / 3600
	var min int64 = (value - hour*3600) / 60
	var second int64 = value - hour*3600 - min*60
	return strconv.FormatInt(hour, 10) + ":" + strconv.FormatInt(min, 10) + ":" + strconv.FormatInt(second, 10)
}

// 时间戳专业字符串
func TimeStampToStr(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

/*
获得多少天后的时间
n 为负数 多少天后
n 为整数 多少天前
2006-01-02 15:04:05
*/
func GetTimeSubDateTime(value string, n int64) string {
	return carbon.Parse(value).SubDays(int(n)).ToDateTimeString()
}

/*
*获得多少天前、后日期 返回 20060102
n 为负数 多少天后
n 为整数 多少天前
*/
func GetTimeSubDate(value string, n int64) string {
	before := GetTimeSubDateTime(value, n)
	beforeDay := carbon.Parse(before).ToDateString()
	beforeDay = strings.ReplaceAll(beforeDay, "-", "")
	return beforeDay
}
