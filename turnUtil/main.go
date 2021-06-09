/*
 * @Author: your name
 * @Date: 2021-06-05 14:56:53
 * @LastEditTime: 2021-06-07 15:59:44
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \goUtils\turnUtil\main.go
 */
package turnUtil

import (
	"strconv"
)

// string转成int：
func StringToInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return -1
	}
	return result
}

// int转成string：
func IntToString(value int) string {
	return strconv.Itoa(value)
}

// int64转成string：
func Int64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}

func StringToInt64(value string) int64 {
	result, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return -1
	}
	return result
}
