## 14-3-2 取得並解析伺服器的 JSON 資料
儘管 HTML、JavaScript 等原始碼適合拿來顯示網頁，但並不適合在機器之間交換資料。  
網路 API 很常使用 JSON 格式資料，因為對人和機器來說，JSON 資料結構良好，能夠輕鬆讀懂。  

以下練習中，要在 Go 程式從伺服器取得結構化的 JSON 資料，並用 json.Unmarshal() 解析成結構形式。  


練習、以 Go HTTP 客戶端存取 JSON 資料：
為模擬 JSON 資料交換，本練習的伺服器和客戶端會由兩支不同的 Go 程式負責。得在專案資料夾下建立兩個子目錄，server、client。  

* 伺服器程式：會在收到客戶端請求後傳回一段簡單的 JSON 資料。
```go
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
```

* 客戶端程式
```go
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
```

* 執行程式：
開起一個主控台，啟動伺服器服務（這伺服器會一直執行到你在主控台 Ctrl+C 中斷它，或關掉主控台為止）  
```shell
Chapter14_GoHTTPClient_20240827/Example14.03/server$ go run server.go 
```

再開啟一個新的主控台，執行客戶端程式。只要伺服器正常運作，它就會從伺服器請求 JSON 資料。
```shell
Go-practice/Chapter14_GoHTTPClient_20240827/Example14.03/client$ go run .
hello world
```
