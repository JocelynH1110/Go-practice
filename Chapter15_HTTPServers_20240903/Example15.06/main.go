package main

import (
	"log"
	"net/http"
)

func main() {
	// 直接將一個匿名函式傳給 HandleFunc()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 把 index.html 當成回應寫入 ResponseWriter
		http.ServeFile(w, r, "./index.html")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
