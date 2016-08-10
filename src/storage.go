/*
	本文件主要是存储上传文件的方法
*/
package main

import (
	"time"
)

//定义上传文件结构体
type File struct {
	Id   string
	Name string
	Size string
	Time string
}

func NewFile(name, size string) File {
	return File{time.Now().Format("20060102150405"), name, size, time.Now().Format("200601021504")}
}

func (file File) Add() {
}
