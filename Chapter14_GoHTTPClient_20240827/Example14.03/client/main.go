package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// 對應 JSON 資料的結構
type messageData struct {
	Message string `json:"message"`
}

func getDataAndReturnResponse() messageData {
	// 對 http://localhost:8080 送出 GET 請求
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	// 從回應主體讀出所有內容的字串
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	message := messageData{}
	// 解析 JSON 字串並資料存入 message（messageData 結構）
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}
	return message
}

func main() {
	data := getDataAndReturnResponse()
	fmt.Println(data.Message)
}
