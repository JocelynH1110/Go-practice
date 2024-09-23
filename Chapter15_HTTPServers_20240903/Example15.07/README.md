## 15-5-2 在伺服器上提供多重靜態資源
前一節學到如何讓使用者存取一個靜態網頁，但若有很多個網頁呢？  
靜態資源不只有 HTML 檔，還包括 JavaScript、CSS 檔、圖片等。要如從伺服器讓網站和使用者存取這些檔案？下面會用幾個 CSS 檔來示範。  

從伺服器提供靜態檔案存取，並將模板分散在不同外部檔案中，通常是將專案問題切割成不同區塊的好辦法，使專案易於管理。  


* 要在 HTML 網頁加入 CSS 樣式檔，可在 <head></head> 之間加入這個標籤：
```html
<link rel="stylesheet" href="myfiles.css">
```
> 這會把名為 myfiles.css 的 CSS 檔嵌入 HTML 網頁，並套用該 CSS 指定的樣式。


* 若檔案系統中要傳回的檔案很多，Go 語言提供了一個函式：
```go
http.FileServer(http.Dir("./public"))
```
> http.FileServer()：會開一個檔案伺服器，而檔案所在的資料夾則是用 http.Dir() 函式取得。
> 上面函式的例子會讓伺服器的 /public 子資料夾下的檔案都能被外界讀取。 如、`http://localhoat:8080/public/myfiles.css` 

檔案以一對一方式傳回：
```go
func ServeFile(w ResponseWriter, r *Request, name string)
```


* 當不想被外界看到伺服器所在機器的目錄名稱時，可以讓檔案伺服器對外提供一個不同的路徑：
```go
http.StripPrefix("/statics/",http.FileServer(http.Dir("./public")))
```
> StripPrefix() 函式：會將請求檔案的 URI 當中的 「/static/」置換成「./public」，並連同檔名傳給檔案伺服器，它會在 ./public 尋找這個檔案。
> 若想沿用原始資料夾名稱，就可以不使用 http.StripPrefix()。


練習、對網頁和使用者提供 CSS 檔  
要展示一個網頁，網頁會引用一些外部 CSS 檔，而這些檔案將透過檔案伺服器來提供。此外檔案在伺服器本地的檔案系統位於 /public 資料夾下。但在伺服器上則透過 /statics 路徑來存取。   

1. 首先建立 HTML 網頁，並用 <link...> 參照到 3 個 CSS 檔:
```html
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <title>Welcome</title>
  <link rel="stylesheet" href="/statics/body.css">
  <link rel="stylesheet" href="/statics/header.css">
  <link rel="stylesheet" href="/statics/text.css">
</head>

<body>
  <h1>Hello</h1>
  <p>Que pasa</p>
</body>

</html>
```

2. 再建立 /public 子資料夾，並撰寫三個 CSS 檔：
* body.css
```
body{
  background-color:orange;
}
```
* header.css
```
h1{
  color:navy;
}
```
* text.css
```
p{
  color:pink;
}
```

3. 最後是伺服器程式 main.go
```go
package main

import (
	"log"
	"net/http"
)

func main() {
	// 對任何路徑提供 index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	// 將 /statics 路徑對應到本地的 /public 資料夾
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("./public"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
