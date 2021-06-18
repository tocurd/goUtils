/*
 * @Author: your name
 * @Date: 2021-06-16 11:22:45
 * @LastEditTime: 2021-06-18 14:42:42
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \webServeri:\project\goUtils\cmdUtil\main.go
 */
package cmdUtil

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"
	"strings"

	"github.com/Tocurd/goUtils/turnUtil"
	"github.com/juju/errors"
)

func Cmd(name string, cmd ...string) (string, error) {

	var out bytes.Buffer
	var outErr bytes.Buffer

	command := exec.Command(name, cmd...)

	command.Stdout = &out
	command.Stderr = &outErr

	err := command.Start()
	if err != nil {
		return outErr.String(), errors.Trace(err)
	}

	err = command.Wait()
	if err != nil {
		return outErr.String(), errors.Trace(err)
	}

	return out.String(), errors.Trace(err)
}

/**
 * @description: 运行CMD
 * @param {string} name
 * @param {...string} cmd
 * @return {*}
 */
func RunCMD(name string, cmd ...string) (string, error) {
	command := exec.Command(name, cmd...)
	stdout, err := command.StdoutPipe() // 获取命令输出内容
	if err != nil {
		return "", errors.Trace(err)
	}
	if err := command.Start(); err != nil { //开始执行命令
		return "", errors.Trace(err)
	}

	useBufferIO := false
	if !useBufferIO {
		var outputBuff bytes.Buffer
		for {
			tempoutput := make([]byte, 256)
			n, err := stdout.Read(tempoutput)
			if err != nil {
				if err == io.EOF { //读取到内容的最后位置
					break
				}
				return "", err
			}
			if n > 0 {
				outputBuff.Write(tempoutput[:n])
			}
		}

		text := turnUtil.ConvertToString(outputBuff.String(), "gbk", "utf8")
		return strings.Trim(text, "\r\n"), err
	} else {
		outputbuff := bufio.NewReader(stdout)
		output, _, err := outputbuff.ReadLine()
		if err != nil {
			return "", err
		}

		text := turnUtil.ConvertToString(string(output), "gbk", "utf8")
		return strings.Trim(text, "\r\n"), nil
	}
}
