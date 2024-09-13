## 15-2-2 簡單的 routing （路由）控制
前一節的伺服器只會回應一個訊息而已，就算客戶端在伺服器路徑下加入不同的子路徑，回應也永遠一樣。  

然而，不同的子路徑可能代表不同的意義。  

好比若向讓伺服器程式顯示一本線上電子書，使用者可用不同的 URL 來存取不同頁數：
```
http://localhost:8080           // 首頁
http://localhost:8080/content   // 目錄頁
http://localhost:8080/page1     // 第一頁
```


* 為了讓伺服器對不同的路徑做出不同回應，得在伺服器套用簡單的 routing（路由）控制，辦法是呼叫 http 套件的 HandleFunc() 函式：
```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```
> pattern：是要處理的子路徑。
> handler：是子路徑被請求時要呼叫的函式，它必須有一個 ResponseWriter 和一個 Request，和前面的 ServeHTTP() 方法相同。


練習、讓伺服器處理路徑
```go
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
```
解析：為何用瀏覽器打開後還是沒有變化？
> 因為對 http.HandleFunc() 函式來說，它設定的對象是 http 套件的 DefaultServeMux 結構。
> DefaultMux 是 http 套件預設的 ServeMux 結構，功能跟我們自己定義的 hello 結構一樣。
> 當我們要程式監聽請求時，卻指定使用結構 hello 為請求處理器，那 DefaultServeMux 就不會發揮功用了。
> 解決方式在下一節。


## 15-2-3 修改程式來應付多重路徑請求
解決方式是：
1. STEP 1、統一使用 DefaultServeMux 來處理客戶端請求，並將 hello 結構註冊給 DefaultServeMux 做為請求處理器：
```go
func Handle(pattern string, handler Handler)
```
> pattern：代表請求處理器要負責的路徑（在此指根目錄/）。
> handler：請求處理器結構。

2. STEP 2、當我們用 ListenAndServe() 啟動伺服器時，它的第二個參數必須設為 nil，這樣它才會使用 DefaultServeMux 來監聽請求：
```go
http.ListenAndServe(":8080",nil)
```

改寫前節練習題為：
```go
func main() {
	http.HandleFunc("/page1", servePage1)
    http.Handle("/",hello{})
	log.Fatal(http.ListenAndServe(":8080",nil) 
}
```

但是其他沒有設定的路徑（例、/page2）就仍會呼叫 hello 的 ServeHTTP() 方法，傳回 http://localhost:8080/ 的頁面。  

* 使用自訂的 ServeMux 結構
自訂的 ServeMux 結構，同樣擁有 HandleFunc() 及 Handle() 方法。  
以下是另一次改寫，建立一個稱為 Mux 的結構，並使用它來接收客戶端請求：
```go
func main(){
    mux:=http.NewServeMux()     // 產生一個新的 ServeMux
    mux.HandleFunc("/page1",servePage1)
    mux.Handle("/",hello{})     // 把 hello 連同其 ServeHTTP() 方法註冊給 mux
    log.Fatal(http.ListenAndServe(":8080",mux))     // 用 mux 來監聽請求
}
```
> 以上執行效果與前面完全相同，只是這裡使用的是 mux 結構而不是 http.DefaultServeMux。


### 回顧：請求處理器 http.Handle() vs. 請求處理函式 http.HandleFunc()
兩者雖然都能處理特定路徑請求，接收的參數卻不相同：
* http.Handle() - 接收一個實作 http.Handler 介面的 **結構**。
```go
func Handle(pattern string, handler Handler)
```
* http.HandleFunc() - 接收一個 **函式**。
```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```


兩者到頭來都會呼叫一個擁有 http.ResponseWriter 和 *http.Request 參數的函式來處理請求，差異不大。  

但選擇正確的做法很重要，尤其是在開發複雜的專案。  
1. http.HandleFunc()：當專案簡單、只有幾個靜態頁面，若是為此特地建一個結構就太多了。
2. http.Handle()：若需要設定一些參數、或追蹤某些資料，把這些資料放在一個結構中就更適當。
