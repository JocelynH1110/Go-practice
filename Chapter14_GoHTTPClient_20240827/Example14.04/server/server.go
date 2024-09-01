package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type server struct{}

type messageData struct {
	Message string `json:"message"`
}

// 伺服器服務
func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := messageData{}
	// 解析客戶端請求主體內的 JSON 資料
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(message)

	// 將訊息轉成全大寫
	message.Message = strings.ToUpper(message.Message)

	// 將 message 重新編碼成 JSON 資料
	jsonBytes, _ := json.Marshal(message)
	w.Write(jsonBytes) // 傳回給客戶端
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
