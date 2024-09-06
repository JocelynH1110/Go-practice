package main

import (
	"log"
	"net/http"
)

type hello struct{} // HTTP 請求處理器

// 請求處理器的方法實作
func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>" // 有 HTML 標籤的文字
	w.Write([]byte(msg))          // 寫入回應傳給客戶端
}
func main() {
	// 啟動伺服器
	log.Fatal(http.ListenAndServe(":8080", hello{}))
}
