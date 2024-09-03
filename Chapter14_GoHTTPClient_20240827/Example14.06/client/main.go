package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func getDataWithCustomOptionsAndReturnResponse() string {
	// 建立一個 Get 請求
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 設定預設客戶端的請求逾時時間為 5 秒
	http.DefaultClient.Timeout = time.Second * 5

	// 將授權碼放入授權標頭 Authorization，值為 "superSecretToken"
	req.Header.Set("Authorization", "superSecretToken")
	resp, err := http.DefaultClient.Do(req) // 讓預設客戶端送出請求
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body) // 讀取回應主體
	if err != nil {
		log.Fatal(err)
	}
	return string(data) // 傳回伺服器回應
}

func main() {
	data := getDataWithCustomOptionsAndReturnResponse()
	fmt.Println(data)
}
