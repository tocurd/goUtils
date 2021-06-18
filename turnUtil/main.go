/*
 * @Author: your name
 * @Date: 2021-06-05 14:56:53
 * @LastEditTime: 2021-06-18 14:41:44
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \goUtils\turnUtil\main.go
 */
package turnUtil

import (
	"strconv"

	"github.com/axgle/mahonia"
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

/**
 * @description:字符串转码
 * @param {string} src
 * @param {string} srcCode
 * @param {string} tagCode
 * @return {*}
 */
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
