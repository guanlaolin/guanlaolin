package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	//配置配置文件
	if !IsExitFile("config.json") {
		log.Fatal("未发现配置文件config.json")
	}
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

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("Listening error:", err)
	}
}
