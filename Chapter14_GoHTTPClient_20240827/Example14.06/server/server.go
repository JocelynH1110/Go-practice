package main

import (
	"log"
	"net/http"
	"time"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 讀取授權標頭 Authorization
	auth := r.Header.Get("Authorization")
	if auth != "superSecretToken" { // 若授權碼不符就拒絕授權
		w.WriteHeader(http.StatusUnauthorized) // 回應設為 HTTP code 401
		w.Write([]byte("Authorization token not recognized"))
		return
	}
	// 授權成功，等待兩秒後回應（模擬處理登入）
	time.Sleep(time.Second * 2)
	msg := "Hello Jocelyn" // 傳回通過授權的訊息
	w.Write([]byte(msg))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
