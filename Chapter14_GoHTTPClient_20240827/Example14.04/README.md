# 14-4 用 POST 請求傳送資料給伺服器
除了從伺服器請求資料，也可以傳資料到伺服器上，常用的方式就是透過 POST 請求。  
常見的例子是：登入表格。  
當按下表格的送出紐，該表格就會對某個 URL 送出 POST 請求。接著網路伺服器通常會檢查登入細節是否正確，正確就更新我們的登入狀態，並回應 POST 請求登入成功。  

* POST 請求是如何傳送資料的呢？  
所有 HTTP 訊息（請求及回應）都包含三部份：URL、標頭（header）、主體（body）。
POST 請求會將要送出的資料夾帶在請求主體中，而不是像 GET 請求那樣透過 URI 參數。


## 14-4-1 送出 POST 請求並接收回應
Go 語言要送出 POST 請求，可使用 http.Post() 函式：
```go
func Post(url, contentType string, body io.Reader) (resp *Response, err error)
```
> url：即請求的網址。
> contentType：為請求標頭中 Content-Type 欄位要指定的內容類型，以一般文字或網頁資料來說就是 "text/html"。 
> body：是 io.Reader 介面型別，即要讓 POST 請求夾帶的資料。


練習、使用 Go HTTP 客戶端會網路伺服器傳送 POST 請求：
要讓客戶端對伺服器送出一個 POST 請求，當中包括一個 JSON 字串。  
然後伺服器會將字串中的 message 訊息轉成全大寫並傳回來。

* 伺服器程式：會實作一個基本的網路伺服器，接收 POST 請求的 JSON 資料並將回覆給客戶端。
```go
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
```
解析：
> json.NewDecoder() 傳回一個 Decoder 結構，並使用後者 Decode() 來解析 JSON 資料。
> NewDecoder() 可接收一個 io.Reader() 介面型別，就不必再用 io.ReadAll() 先將它轉成字串了。


* 客戶端程式：
```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type messageData struct {
	Message string `json:"message"`
}

func postDataAndReturnResponse(msg messageData) messageData {
	jsonBytes, _ := json.Marshal(msg)    // 將要傳的結構編碼成 JSON 資料
	buffer := bytes.NewBuffer(jsonBytes) // 將 JSON 字串轉成 bytes.Buffer

	// 送出 POST 請求和資料，標頭為 application/json ,並接收回應
	r, err := http.Post("http://localhost:8080", "application/json", buffer)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	message := messageData{}
	// 解碼伺服器回應的 JSON 資料
	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Fatal(err)
	}
	return message
}
func main() {
	msg := messageData{Message: "Hoho server!"} // 要傳給伺服器的訊息
	data := postDataAndReturnResponse(msg)      // 接收伺服器傳回的訊息

	fmt.Println(data.Message)
}
```


* 執行程式：
1. 先執行 server
```shell
Chapter14_GoHTTPClient_20240827/Example14.04/server$ go run .
```

2. 在執行 client 目錄下的 main.go
```shell
Go-practice/Chapter14_GoHTTPClient_20240827/Example14.04/client$ go run .
HOHO SERVER!
```
> 傳的字都變成大寫了

3. 在執行 server 的主控台中，也會看到伺服器印出從客戶端收到的資料：
```shell
Go-practice/Chapter14_GoHTTPClient_20240827/Example14.04/server$ go run .
2024/08/31 15:29:22 {Hoho server!}
```
