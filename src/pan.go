//pan.guanlaolin.cn Handler
package main

import (
	"net/http"
)

//pan.guanlaolin.cn/
func PanIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//GET方法，向前台传送模板

	}
}
