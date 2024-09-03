# 14-5 在客戶端使用自訂標頭做為請求選項
對任何請求，都可以設定其他標頭，甚至加入自訂得標頭，做為對伺服器的請求選項。  

常見的例子：授權標頭。

當你註冊了一個服務時，會傳回一個授權碼，然後呼叫該服務的 API 時，請求裡都必須包含這個授權碼，好讓伺服器驗證你的身份。  

1. 為了能設定自訂標頭，得使用 http.NewRequest() 產生一個 http.Request 結構：
```go
func NewRequest(method, url string, body io.Reader)(*Request, error)
```
> method：為 HTTP 方法（即 GET、POST 等）。
> url：網址。
> body：請求主體。
> 取得請求物件後，就可以對它指定標頭。

2. 要使用一個 http.Client 結構的 Do() 方法來執行這個請求：
```go
func (c *Client) Do(req *Request) (*Response, error)
```

可以隨意建立 Client 結構，但這練習會用 http 套件的預設客戶端 DefaultClient。  


練習、在 GET 請求加入授權標頭：
打造自己的 HTTP 客戶端，並加入授權標頭—伺服器只有在請求內包含正確得標頭和授權碼時，才會傳回你需要的訊息。  

* 伺服器程式：
```go
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
```
> Authorization 是 HTTP 協定中已存在的標頭欄位。
> http.ResponseWriter 介面的 WriteHeader() 方法能用來設定回應的 HTTP 狀態碼，如果沒呼叫它就會預設傳回 http.StatusOK（200）。

* 客戶端伺服器：
```go
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
```

顯示結果：
```shell
$ go run .
Hello Jocelyn
```
