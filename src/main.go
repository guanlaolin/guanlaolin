package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	PORT = "port"
)

const (
	CONFIG_FILE = "../config.json" //配置文件路径

	HTML_DIR = "../statics/" //html路径
)

var (
	configs   map[string]string             //保存配置文件
	templates map[string]*template.Template //保存静态模板
)

func init() {
	//配置配置文件
	if !IsExitFile(CONFIG_FILE) {
		log.Fatal("未发现配置文件:", CONFIG_FILE)
	}

	f, _ := os.Open(CONFIG_FILE)
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("读取配置文件信息失败：", err)
	}
	if err = json.Unmarshal(bs, &configs); err != nil {
		log.Fatal("解析配置文件失败：", err)
	}
	//log.Print(configs)

	//加载静态模板，缓存到内存以加快访问速度
	files, err := ioutil.ReadDir(HTML_DIR)
	if err
	template.ParseFiles()
}

func main() {
	r := mux.NewRouter()

	//guanlaolin.cn
	main := r.Host("www.guanlaolin.cn").Subrouter()
	for url, handler := range main_urls {
		main.HandleFunc(url, handler)
	}

	//pan.guanlaolin.cn
	pan := r.Host("pan.guanlaolin.cn").Subrouter()
	for url, handler := range pan_urls {
		pan.HandleFunc(url, handler)
	}

	if err := http.ListenAndServe(configs[PORT], r); err != nil {
		log.Fatal("Listening ", configs[PORT], " error:", err)
	}
}
