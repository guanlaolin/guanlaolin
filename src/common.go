/*
	本文件主要是一些通用函数
*/
package main

import (
	"os"
)

//判断文件是否存在，如果文件存在返回true
func IsExitFile(file string) bool {
	f, err := os.Open(file)

	defer f.Close()

	if os.IsNotExist(err) {
		return false
	}

	return true
}
