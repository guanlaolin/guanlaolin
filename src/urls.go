/**
本文件是网站路由映射
*/
package main

import (
	"net/http"
)

//guanlaolin.cn路由映射
var main_urls = map[string]func(w http.ResponseWriter, r *http.Request){
	"/": IndexHandler,
}

//pan.guanlaolin.cn路由映射
var pan_urls = map[string]func(w http.ResponseWriter, r *http.Request){
	"/":         PanIndexHandler,
	"/upload":   PanUploadHandler,
	"/list":     PanListHandler,
	"/download": PanDownloadHandler,
}
