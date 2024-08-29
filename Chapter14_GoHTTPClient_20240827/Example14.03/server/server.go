package main

import (
	"log"
	"net/http"
)

type server struct{} // 伺服器結構

// 收到請求時要執行的伺服器服務
func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := `{"message":"hello world"}` // 要傳回給客戶端的 JSON 資料
	w.Write([]byte(msg))               // 將 JSON 字串寫入回應主體
}

func main() {
	// 啟動本地伺服器，監聽 port 8080
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
