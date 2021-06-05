/*
 * @Author: your name
 * @Date: 2021-06-04 22:28:48
 * @LastEditTime: 2021-06-05 10:40:50
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edits
 * @FilePath: \webServeri:\project\goUtils\time.go
 */
package goUtils

import "time"

// 当前时间
func GetTime() time.Time {
	return time.Now()
}

// 格式化为:20060102150405
func GetTimeString(t time.Time, timeFormat string) string {
	return t.Format(timeFormat)
}

// 格式化为:2006-01-02 15:04:05
func GetNormalTimeString(t time.Time, normalTimeFormat string) string {
	return t.Format(normalTimeFormat)
}

// 转为时间戳->秒数
func GetTimeUnix(t time.Time) int64 {
	return t.Unix()
}

// 转为时间戳->毫秒数
func GetTimeMills(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// 时间戳转时间
func GetTimeByInt(t1 int64) time.Time {
	return time.Unix(t1, 0)
}

// 字符串转时间
func GetTimeByString(timestring string, timeFormat string) (time.Time, error) {
	if timestring == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(timeFormat, timestring, time.Local)
}

// 标准字符串转时间
func GetTimeByNormalString(timestring string, normalTimeFormat string) (time.Time, error) {
	if timestring == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(normalTimeFormat, timestring, time.Local)
}

// 比较两个时间大小
func CompareTime(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

// n小时后的时间字符串
func GetNextHourTime(s string, n int64, timeFormat string) string {
	t2, _ := time.ParseInLocation(timeFormat, s, time.Local)
	t1 := t2.Add(time.Hour * time.Duration(n))
	return GetTimeString(t1, timeFormat)
}

// 计算俩个时间差多少小时
func GetHourDiffer(start_time, end_time string, timeFormat string) float32 {
	var hour float32
	t1, _ := time.ParseInLocation(timeFormat, start_time, time.Local)
	t2, err := time.ParseInLocation(timeFormat, end_time, time.Local)
	if err == nil && CompareTime(t1, t2) {
		diff := GetTimeUnix(t2) - GetTimeUnix(t1)
		hour = float32(diff) / 3600
		return hour
	}
	return hour
}

// 判断当前时间是否是整点
func Checkhours() bool {
	_, m, s := GetTime().Clock()
	if m == s && m == 0 && s == 0 {
		return true
	}
	return false
}
