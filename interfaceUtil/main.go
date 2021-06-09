/*
 * @Author: your name
 * @Date: 2021-06-09 17:49:48
 * @LastEditTime: 2021-06-09 17:54:46
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \webServeri:\project\goUtils\interfaceUtil\main.go
 */
package interfaceUtil

import (
	"errors"
	"reflect"
)

/**
 * @description: 获取interface struct中的所有的字段名
 * @param {interface{}} value
 * @return {*}
 */
func Fields(value interface{}) ([]string, error) {
	typeof := reflect.TypeOf(value)
	if typeof.Kind() == reflect.Ptr {
		typeof = typeof.Elem()
	}

	var fields = make([]string, 0)
	// 强制要求必须输入结构体
	if typeof.Kind() != reflect.Struct {
		return fields, errors.New("check type error not struct")
	}

	for index := 0; index < typeof.NumField(); index++ {
		jsonKey := typeof.Field(index).Tag.Get("json")
		if jsonKey != "" {
			fields = append(fields, jsonKey)
		} else {
			fields = append(fields, typeof.Field(index).Name)
		}
	}

	return fields, nil
}
