# 14-3 對伺服器傳送 GET 請求
GET 是 HTTP 請求中最常見的一種，它使用的 URL 會描述資源所在的位址，並藉由查詢參數來附帶額外的資訊。  

GET 請求的 URL 能拆成幾個部份：  
`http://example.com/downloads?filter=latest&os=windows`  

1. 協定（protocol）：http。描述客戶端如何跟伺服器連線，最常見的協定為 HTTP、HTTPS。
2. 主機名稱（hostname）：example。要連上的伺服器的位址。
3. URI：全名為 **統一資源識別碼（Uniform Resource Identifier）**，為伺服器資源的所在路徑。
4. 查詢參數（query parameters）：參數與 URI 是以問號分開的，好讓伺服器能解析出參數。各參數間用＆號分隔。

## 14-3-1 使用 http.GET() 發送 GET 請求  
若想在 Go 語言對一個 URL 送出 GET 請求，並接收伺服器的回應，  
可使用 http.Get()：
```go
func Get(url string)(resp *Response, err error)
```
> 傳回的 http.Response 結構中，屬性 Body （即 request body，回應主體）會包含傳回的內容，事後應該用 Body.Close() 關閉它。
> 且它符合 io.Writer 介面的規範，可用 io.ReadAll 來讀取內容。
> Response 的屬性 StatusCode 則是請求的 HTTP 狀態碼（整數），通常 200 代表請求成功。
> 即使狀態碼不為 2xx ，http.Get() 也不會傳回錯誤。得檢查 StatusCode 才能知道請求是否成功。  


練習、用 Go HTTP 客戶端對網路伺服器傳送 GET 請求
```go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getDataAndReturnResponse() string {
	// 送出 get 請求和取得回應
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() // 結束時，釋放 r 佔用的連線資源

	if resp.StatusCode != http.StatusOK { // 檢查 HTTP 狀態碼是否不為 200（ok）
		log.Fatal(resp.Status) // 狀態碼不為 200 的話，用 Status 屬性印出完整狀態碼描述
	}

	data, err := io.ReadAll(resp.Body) // 讀取回應主題的所有內容
	if err != nil {
		log.Fatal(err)
	}
	return string(data) // 傳回回應內容
}
func main() {
	data := getDataAndReturnResponse()
	fmt.Print(data)
}
```
結果顯示：
```shell
    <!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="zh-TW"><head><meta content="text/html; charset=UTF-8" http-equiv="Content-Type"><meta content="/images/branding/googleg/1x/googleg_standard_color_128dp.png" itemprop="image"><title>Google</title><script nonce="_VEXIJnGWzMLPfKoksvq_w">(function(){var _g={kEI:'fuXOZo_MM8uSvr0P-cLnwQc',kEXPI:'0,3700331,618,432,7,90,538567,2872,2891,73050,16105,163859,2,39761,6699,41946,57737,2,2,1,10957,15675,8155,....以下略
```
解析：  
> 以上對 google 送出 GET 請求，並將伺服器的回應印在主控台。  
> 乍看像是亂碼，但若存在 HTML 檔和用瀏覽器打開，會出現 google 首頁。  
> 這就是 **從伺服器取得並解讀結構化的資料、然後顯示成網頁**。
