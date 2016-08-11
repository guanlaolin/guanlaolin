//pan.guanlaolin.cn Handler
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

const (
	UPLOAD_DIR = "../publics/"

	INFO_FILE = UPLOAD_DIR + "storage"
)

//文件信息
type Info struct {
	Id   string `json:"url"`
	Name string `json:"name"`
	Time string `json:"time"`
}

func NewInfo(id, name, time string) Info {
	return Info{id, name, time}
}

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
		//读取INFO_FILE文件获取文件信息
		f, err := os.Open(INFO_FILE)
		if err != nil {
			log.Println("Open file error:", err)
			http.Error(w, "获取文件列表失败，请重试", http.StatusServiceUnavailable)
			return
		} //if
		defer f.Close()
		file, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println("Copy file error:", err)
			http.Error(w, "获取文件列表失败，请重试", http.StatusServiceUnavailable)
			return
		} //if
		infoArr := bytes.Split(file, []byte("\n"))
		//log.Printf("%s", infoArr)
		var tempInfoArr []Info

		for i, info := range infoArr {
			infoSli := bytes.Split(info, []byte("\t"))
			//因为infoArr最后一行为空，必须减1
			if i == len(infoArr)-1 {
				continue
			}
			//log.Print(string(infoSli[0]), string(infoSli[1]), string(infoSli[2]))
			tempInfoArr = append(tempInfoArr,
				//info[0]:id;info[1]name;info[3]time
				NewInfo(string(infoSli[0]),
					string(infoSli[1]),
					string(infoSli[3])))
		}
		infoJson, err := json.Marshal(tempInfoArr)
		if err != nil {
			log.Println("Copy file error:", err)
			http.Error(w, "获取文件列表失败，请重试", http.StatusServiceUnavailable)
			return
		} //if
		_, err = w.Write(infoJson)
		if err != nil {
			log.Println("Copy file error:", err)
			http.Error(w, "获取文件列表失败，请重试", http.StatusServiceUnavailable)
			return
		} //if

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
		file, err := os.OpenFile(INFO_FILE, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
		if err != nil {
			log.Print("Open file storage.txt error:", err)
			http.Error(w, "上传文件失败", http.StatusServiceUnavailable)
			os.Remove(id)
			return
		}
		defer file.Close()

		//id+文件名+类型+上传时间
		_, err = file.WriteString(id + "\t" + h.Filename + "\t" + path.Ext(h.Filename) + "\t" + time.Now().Format("200601021504") + "\n")
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

//pan.guanlaolin.cn/download get方法下载文件
func PanDownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.FormValue("id")
		var name string //文件实际名字
		//var mime string //文件类型
		if id == "" {
			http.Error(w, "请求id为空", http.StatusBadRequest)
			return
		}
		//读取INFO_FILE文件获取文件信息
		f, err := os.Open(INFO_FILE)
		if err != nil {
			log.Println("Open file error:", err)
			http.Error(w, "下载文件失败，请重试", http.StatusServiceUnavailable)
			return
		} //if
		defer f.Close()
		file, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println("Copy file error:", err)
			http.Error(w, "下载文件失败，请重试", http.StatusServiceUnavailable)
			return
		} //if
		infoArr := bytes.Split(file, []byte("\n"))
		//log.Printf("%s", infoArr)

		for i, info := range infoArr {
			infoSli := bytes.Split(info, []byte("\t"))
			//因为infoArr最后一行为空，必须减1
			if i == len(infoArr)-1 {
				continue
			}

			//找到id对应的文件
			if string(infoSli[0]) == id {
				name = string(infoSli[1])
				//mime = string(infoSli[2])
			}
		} //for

		if name == "" {
			http.Error(w, "请求id有误，请重试", http.StatusBadRequest)
		}
		w.Header().Set("Content-Disposition", "attachment;filename="+name)
		path := UPLOAD_DIR + id
		http.ServeFile(w, r, path)

	} else {
		http.Error(w, "未实现方法", http.StatusFound)
	}

}
