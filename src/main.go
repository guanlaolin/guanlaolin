package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
)

const (
	PORT = "port"

	EXT_HTML = ".html"
)

const (
	CONFIG_FILE = "../config.json" //配置文件路径

	HTML_DIR = "../statics/"     //html路径
	CSS_DIR  = HTML_DIR + "css/" //css路径
	JS_DIR   = HTML_DIR + "js/"  //js路径
)

var (
	configs   = make(map[string]string)             //保存配置文件
	templates = make(map[string]*template.Template) //保存静态模板
)

//系统初始化
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
	if err != nil {
		log.Print("读取模板目录失败：", err)
	}

	for _, file := range files {
		if path.Ext(file.Name()) != EXT_HTML {
			continue
		}
		tmpl, err := template.ParseFiles(HTML_DIR + file.Name())
		if err != nil {
			log.Print("解析模板：", file.Name(), "失败：", err)
		}

		templates[file.Name()] = tmpl
	}
	//log.Print(templates)
}

func main() {

	r := mux.NewRouter()

	//静态文件处理
	http.Handle("/css/",
		http.StripPrefix("/css/",
			http.FileServer(http.Dir(CSS_DIR))))
	http.Handle("/js/",
		http.StripPrefix("/js/",
			http.FileServer(http.Dir(JS_DIR))))
	//用handler代替
	//	http.Handle("/publics/",
	//		http.StripPrefix("/publics/",
	//			http.FileServer(http.Dir(UPLOAD_DIR))))

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

	http.Handle("/", r)

	if err := http.ListenAndServe(configs[PORT], nil); err != nil {
		log.Fatal("Listening ", configs[PORT], " error:", err)
	}
}
