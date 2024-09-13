package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	v1 := r.URL.Query()    // 讀取查詢字串
	name, ok := v1["name"] // 讀取參數 name
	if !ok {               // 若查無參數
		w.WriteHeader(http.StatusBadRequest) // 回應 http 404 （bad request）
		w.Write([]byte("<h1>Missing name</h1>"))
		return
	}
	// 在網頁產生針對使用者的歡迎訊息
	w.Write([]byte(fmt.Sprintf("<h1>Hello %s</h1>", strings.Join(name, ","))))
}
func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
