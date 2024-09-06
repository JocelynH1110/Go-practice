# 15-2 打造最基本的伺服器
我們所能寫出的最基本 HTTP 伺服器叫做 Hello World 伺服器：  
當使用者存取伺服器的位址（如 http://localhost:8080）時，只會傳回一句文字 Hello World，沒有其他功能。

下面會撰寫一個伺服器，如何在普通網頁連覽器（客戶端）中顯示結果。

## 15-2-1 使用 HTTP 請求處理器（handler）
handle（請求處理器）：為了應付 HTTP 請求，會需要寫一個功能來處理（handle）請求。這功能稱為一個 handle。  

在 Go 語言有幾種方式能寫請求處理器，其中一個是實作 http 套件的 Handler 介面。  
這介面只有一個方法 ServeHTTP() ：
```go
ServeHTTP(w http.ResponseWriter,r *http.Request)
```
> 會接收一個 http.Request 型別，也就是來自客戶端的請求。能從他的標頭和主體讀取資料。
> http.ResponseWriter 型別：即為要寫入給客戶端的回應。


* 實作 HTTP 請求處理器：
1. 可先建立一個沒內容的空結構，再掛上 ServeHTTP() 方法：
```go
type myHandler struct{}

func (h MyHandler) ServeHTTP(w http.RespnseWriter,r *http.Request){}
```

2. 上面的 myHandler 便成為了合法的 HTTP 請求處理器。於是就能用 http 套件的 ListenAndServe() 函式來監聽指定的 TCP 位址：  
```go
http.ListenAndServe(":8080",myHandler{})
```
> 此舉等於是啟動伺服器、監聽 http://localhost:8080 和等待客戶端送出請求，若收到就會呼叫 myHandler.ServeHTTP()。

3. ListenAndServe() 函式可能會傳回一個 error，我們在此情況下很可能會希望回報錯誤並中止程式。  
故常見的做法為---把這函式包在 log.Fatal() 內：
```go
log.Fatal(http.ListenAndServe(":8080",myHandler{}))
```

練習、建立 Hello World 伺服器
```go
package main

import (
	"log"
	"net/http"
)

type hello struct{} // HTTP 請求處理器

// 請求處理器的方法實作
func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>" // 有 HTML 標籤的文字
	w.Write([]byte(msg))          // 寫入回應傳給客戶端
}
func main() {
	// 啟動伺服器
	log.Fatal(http.ListenAndServe(":8080", hello{}))
}
```
> 正常情況下 Go 語言會呼叫 ResponseWriter 的方法 WriteHeader()，並填入 HTTP 狀態碼 http.StatusOK（即 200）。  

顯示結果：
1. 至新主控台執行
```shell
Example15.02$ go run .
```

2. 執行後伺服器不會有任何訊息。要到網頁瀏覽器，輸入 http://localhost:8080

3. 倘若嘗試改變路盡，比如輸入 http://localhost:8080/page1 ，還是會收到一樣的訊息。
> 中止伺服器：ctrl+c。
> 執行其他範例看到舊的範例畫面：ctrl+F5。可能是瀏覽器的取畫面，可強制更新。
