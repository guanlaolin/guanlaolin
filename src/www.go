//ww.guanlaolin Handler
package main

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates["index.html"].Execute(w, nil)
	}
}
