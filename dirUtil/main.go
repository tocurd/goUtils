/*
 * @Author: your name
 * @Date: 2021-06-09 17:36:31
 * @LastEditTime: 2021-06-09 17:37:15
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \webServeri:\project\goUtils\dirUtil\main.go
 */
package dirUtil

import "os"

func IsDir(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
