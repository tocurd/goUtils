/*
 * @Author: your name
 * @Date: 2021-06-05 14:54:31
 * @LastEditTime: 2021-06-05 14:55:38
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \goUtils\hashUtil\md5.go
 */
package hashUtil

import (
	cryptoMd5 "crypto/md5"
	"fmt"
)

func Md5(value string) string {
	return fmt.Sprintf("%x", cryptoMd5.Sum([]byte(value)))
}
