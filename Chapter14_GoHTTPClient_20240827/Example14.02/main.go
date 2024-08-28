package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getDataAndReturnResponse() string {
	// 送出 get 請求和取得回應
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() // 結束時，釋放 r 佔用的連線資源

	if resp.StatusCode != http.StatusOK { // 檢查 HTTP 狀態碼是否不為 200（ok）
		log.Fatal(resp.Status) // 狀態碼不為 200 的話，用 Status 屬性印出完整狀態碼描述
	}

	data, err := io.ReadAll(resp.Body) // 讀取回應主題的所有內容
	if err != nil {
		log.Fatal(err)
	}
	return string(data) // 傳回回應內容
}
func main() {
	data := getDataAndReturnResponse()
	fmt.Print(data)
}
