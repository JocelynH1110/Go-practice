package main

import (
	"log"
	"net/http"
)

type hello struct{}

// 原本的請求處理方法
func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>"
	w.Write([]byte(msg))
}

// 新的函式，用來處理對路徑 /page1 的請求
func servePage1(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Page 1</h1>"
	w.Write([]byte(msg))
}

func main() {
	// 在客戶端請求路徑 /page1 時，呼叫 servePage1()
	http.HandleFunc("/page1", servePage1)

	// 監聽 localhost:8080 並在需要時呼叫 hello.ServeHTTP()
	log.Fatal(http.ListenAndServe(":8080", hello{}))
}
