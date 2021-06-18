/*
 * @Author: your name
 * @Date: 2021-06-18 10:35:34
 * @LastEditTime: 2021-06-18 10:36:28
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \webServeri:\project\goUtils\pathUtil\main.go
 */
package pathUtil

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// func execPath() (string, error) {
// 	file, err := exec.LookPath(os.Args[0])
// 	if err != nil {
// 		return "", err
// 	}
// 	_, err = filepath.Abs(file)
// 	if err != nil {
// 		return "", err
// 	}
// 	return filepath.Abs(file)
// }
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径 filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
