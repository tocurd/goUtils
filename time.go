/*
 * @Author: your name
 * @Date: 2021-06-04 22:28:48
 * @LastEditTime: 2021-06-05 09:01:51
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \webServeri:\project\goUtils\time.go
 */
package goUtils

import "time"

type TimeUtil struct {
	TimeFormat       string "20060102150405"
	NormalTimeFormat string "2006-01-02 15:04:05"
}

// 当前时间
func (utilTime *TimeUtil) GetTime() time.Time {
	return time.Now()
}

// 格式化为:20060102150405
func (utilTime *TimeUtil) GetTimeString(t time.Time) string {
	return t.Format(utilTime.TimeFormat)
}

// 格式化为:2006-01-02 15:04:05
func (utilTime *TimeUtil) GetNormalTimeString(t time.Time) string {
	return t.Format(utilTime.NormalTimeFormat)
}

// 转为时间戳->秒数
func (utilTime *TimeUtil) GetTimeUnix(t time.Time) int64 {
	return t.Unix()
}

// 转为时间戳->毫秒数
func (utilTime *TimeUtil) GetTimeMills(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// 时间戳转时间
func (utilTime *TimeUtil) GetTimeByInt(t1 int64) time.Time {
	return time.Unix(t1, 0)
}

// 字符串转时间
func (utilTime *TimeUtil) GetTimeByString(timestring string) (time.Time, error) {
	if timestring == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(utilTime.TimeFormat, timestring, time.Local)
}

// 标准字符串转时间
func (utilTime *TimeUtil) GetTimeByNormalString(timestring string) (time.Time, error) {
	if timestring == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(utilTime.NormalTimeFormat, timestring, time.Local)
}

// 比较两个时间大小
func (utilTime *TimeUtil) CompareTime(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

// n小时后的时间字符串
func (utilTime *TimeUtil) GetNextHourTime(s string, n int64) string {
	t2, _ := time.ParseInLocation(utilTime.TimeFormat, s, time.Local)
	t1 := t2.Add(time.Hour * time.Duration(n))
	return utilTime.GetTimeString(t1)
}

// 计算俩个时间差多少小时
func (utilTime *TimeUtil) GetHourDiffer(start_time, end_time string) float32 {
	var hour float32
	t1, _ := time.ParseInLocation(utilTime.TimeFormat, start_time, time.Local)
	t2, err := time.ParseInLocation(utilTime.TimeFormat, end_time, time.Local)
	if err == nil && utilTime.CompareTime(t1, t2) {
		diff := utilTime.GetTimeUnix(t2) - utilTime.GetTimeUnix(t1)
		hour = float32(diff) / 3600
		return hour
	}
	return hour
}

// 判断当前时间是否是整点
func (utilTime *TimeUtil) Checkhours() bool {
	_, m, s := utilTime.GetTime().Clock()
	if m == s && m == 0 && s == 0 {
		return true
	}
	return false
}
