# 15-3 解讀網址參數來動態產生網頁
HTTP 伺服器可以根據更多的請求細節產生回應，這些細節不僅能用路徑的形式，更可用網址參數傳給伺服器。  
參數的傳遞方式有很多種，最常見的是 **查詢字串（QueryString）**，它包含所謂的查詢參數。  

```go
http://localhost:8080?name=john
```
> ?name=john 這部份就是查詢字串。參數是 name，值被指定為 john。

如果有更多參數，會用 & 來連接：
```go
http://localhost:8080?name=john&age=20
```

查詢字串通常會搭配 GET 請求使用，因為 POST 請求一般會透過請求主體而不是查詢參數來傳遞資料。  
若想解讀客戶端請求中 URL 夾帶的參數，要透過 http.Request 結構 URL 屬性的 Query() 方法。

練習、顯示個人化的歡迎訊息
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	v1 := r.URL.Query()    // 讀取查詢字串
	name, ok := v1["name"] // 讀取參數 name
	if !ok {               // 若查無參數
		w.WriteHeader(http.StatusBadRequest) // 回應 http 404 （bad request）
		w.Write([]byte("<h1>Missing name</h1>"))
		return
	}
	// 在網頁產生針對使用者的歡迎訊息
	w.Write([]byte(fmt.Sprintf("<h1>Hello %s</h1>", strings.Join(name, ","))))
}
func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
> r.URL.Query() 會傳回 map[string][]string 型別，其鍵為參數名稱，對應值為參數值。  
> 參數值是字串切片，因使用者有可能用 ?name=名字1,名字2... 的方式輸入不只一個姓名。
> strings.Join() 將 name 切片內的所有元素連接起來，並以逗號連接成單一一個字串。如此一來，即使輸入多重人名，程式也能解讀並正確顯示出來。
