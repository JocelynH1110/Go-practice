# 15-5 使用靜態網頁資源
學到目前為止，伺服器傳回的東西都是在程式裡定義好的結果。若想更改訊息，就得修改程式碼、重新編譯、重新執行伺服器。  
要是把這程式賣給別人，要讓他們能自行修改訊息，就得提供他們程式碼。  
開放原始碼或許不是壞事，但用這種方式散播應用程式絕非理想之道。  

解決方法：在伺服器使用靜態檔案。  
 
他們會被程式當成外部檔案載入。比如上一節的模板就可做成檔案，模板只是文字，可以把它寫在檔案內而不是程式碼中。  
其他常見的靜態資源還包括：網頁圖片、CSS 樣式檔、JavaScript 指令檔等等。  


## 15-5-1 讀取靜態 HTML 網頁
下面練習會講解怎麼在伺服器載入特定目錄下的特定靜態檔案。  

為了把一個靜態檔案當成 HTTP 回應傳送給客戶端，得在請求處理器的 ServeHTTP() 方法或請求處理函式接收的函式內呼叫 http.ServeFile()：
```go
func ServeFile(w ResponseWriter, r *Request, name string)
```

練習、使用靜態網頁的 Hello World 伺服器：
改用靜態 HTML 網頁做為輸出。只需使用一個請求處理函式，適用於該伺服器的所有 URL 路徑。

1. 首先在專案下建立一個 HTML 文字檔 index.html：  

```html
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <title>Welcome</title>
</head>

<body>
  <h1>Hello World Haha</h1>
</body>

</html>
```

2. 撰寫 main.go：   

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	// 直接將一個匿名函式傳給 HandleFunc()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 把 index.html 當成回應寫入 ResponseWriter
		http.ServeFile(w, r, "./index.html")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

解析：
打開網頁瀏覽器會出現 Hello World，然後在不關閉伺服器程式的狀況下修改 index.html 內要顯示在網頁瀏覽器上的內容。  
再重新整理瀏覽器，會發現即使伺服器仍在運作，照樣能載入修改過的 HTML 檔。  

換言之，當把靜態資源跟伺服器程式邏輯切割開來時，就能在不關閉伺服器的情況下更新網頁。
