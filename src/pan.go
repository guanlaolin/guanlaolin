//pan.guanlaolin.cn Handler
package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	UPLOAD_DIR = "../publics/"
)

//pan.guanlaolin.cn/
func PanIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//GET方法，向前台传送模板
		templates["list.html"].Execute(w, nil)
	} else {
		http.Error(w, "未实现方法", http.StatusNotImplemented)
	}
}

//pan.guanlaolin.cn/list get方法返回json
func PanListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

	} else {
		http.Error(w, "未实现方法", http.StatusNotImplemented)
	}
}

//pan.guanlaolin.cn/upload post方法实现文件上传
func PanUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("data")
		if err != nil {
			//返回503错误
			http.Error(w, "上传文件失败", http.StatusServiceUnavailable)
			return
		}
		defer f.Close()

		id := time.Now().Format("20060102150405")
		temp, err := os.Create(UPLOAD_DIR + id)
		if err != nil {
			log.Print(" Create file ", id, " error:", err)
			http.Error(w, "上传文件失败", http.StatusServiceUnavailable)
			return
		}
		defer temp.Close()

		if _, err = io.Copy(temp, f); err != nil {
			log.Println("Copy file error:", err)
			http.Error(w, "上传文件失败", http.StatusServiceUnavailable)
			return
		}

		//将上传文件存储到文件信息表
		file, err := os.OpenFile("storage.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
		if err != nil {
			log.Print("Open file storage.txt error:", err)
			http.Error(w, "上传文件失败", http.StatusServiceUnavailable)
			os.Remove(id)
			return
		}
		defer file.Close()

		_, err = file.WriteString(id + "\t" + h.Filename + "\t" + time.Now().Format("200601021504") + "\r\n")
		if err != nil {
			log.Print("Write file inform error:", err)
			http.Error(w, "上传文件失败", http.StatusServiceUnavailable)
			os.Remove(id)
			return
		}
		//上传成功，跳转到pan.guanlaolin.cn
		http.Redirect(w, r, "/", http.StatusFound)

	} else {
		http.Error(w, "未实现方法", http.StatusNotImplemented)
	}
}
