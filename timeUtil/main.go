/*
 * @Author: your name
 * @Date: 2021-06-04 22:28:48
 * @LastEditTime: 2021-06-05 11:03:47
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edits
 * @FilePath: \webServeri:\project\goUtils\time.go
 */
package timeUtil

import "time"

/**
 * @description: 获取今日0点的时间戳
 * @param {*}
 * @return {*}
 */
func GetNowDay() (time.Time, error) {
	nowDay := time.Now().Format("2016-01-26")
	return GetTimeByString(nowDay, "2016-01-26")
}

/**
 * @description: 当前时间
 * @param {*}
 * @return {*}
 */
func GetTime() time.Time {
	return time.Now()
}

/**
 * @description: 将时间转换成指定字符串
 * @param {time.Time} t 时间
 * @param {string} timeFormat 时间格式
 * @return {*}
 */
func GetTimeString(t time.Time, timeFormat string) string {
	return t.Format(timeFormat)
}

/**
 * @description: 将时间转换成时间戳
 * @param {time.Time} t
 * @return {*}
 */
func GetTimeUnix(t time.Time) int64 {
	return t.Unix()
}

/**
 * @description: 将秒级时间戳转换成毫秒
 * @param {time.Time} t
 * @return {*}
 */
func GetTimeMills(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

/**
 * @description: 将时间戳转换成时间
 * @param {int64} t1
 * @return {*}
 */
func GetTimeByInt(t1 int64) time.Time {
	return time.Unix(t1, 0)
}

/**
 * @description: 将字符串转换成时间格式
 * @param {string} timestring
 * @param {string} timeFormat 字符串的格式
 * @return {*}
 */
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

/**
 * @description: 比较两个时间大小
 * @param {*} t1
 * @param {time.Time} t2
 * @return {*}
 */
func CompareTime(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

/**
 * @description: 获取基于当前时间N小时候的时间字符串
 * @param {string} s
 * @param {int64} n
 * @param {string} timeFormat
 * @return {*}
 */
func GetNextHourTime(s string, n int64, timeFormat string) string {
	t2, _ := time.ParseInLocation(timeFormat, s, time.Local)
	t1 := t2.Add(time.Hour * time.Duration(n))
	return GetTimeString(t1, timeFormat)
}

/**
 * @description: 计算俩个时间差多少小时
 * @param {*} start_time
 * @param {string} end_time
 * @param {string} timeFormat
 * @return {*}
 */
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

/**
 * @description: 判断当前时间是否是整点
 * @param {*}
 * @return {*}
 */
func Checkhours() bool {
	_, m, s := GetTime().Clock()
	if m == s && m == 0 && s == 0 {
		return true
	}
	return false
}
