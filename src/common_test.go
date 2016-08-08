package main

import (
	"log"
)

import (
	"testing"
)

//测试IsExistFile函数
func TestIsExistFile(T *testing.T) {
	if IsExitFile("testfile") {
		log.Print("false")
	} else if IsExitFile("common.go") {
		log.Print("true")
	}
}
