# 11-3 將 Go 結構編碼為 JSON
將存在結構裡的資料編碼（marshal）成 JSON 格式。  

何時使用？  
將資料從檔案、資料庫讀出來和轉成 JSON 格式，以便透過網路傳給請求者，或為了將資料寫入 NoSQL 資料庫，得先將它轉成 JSON 才行。
## 11-3-1 Marshal()
用 encoding/json 套件的 Marshal() 函式來編碼：
```go
func Marshal(v interface{})([]byte, error)
```
> * 參數 v：需要編碼成 JSON 格式的原始資料，通常是個結構。
> * 傳回值：會傳回 JSON 字串（[]byte 切片）以及 error 值，如編碼失敗 error 就不為 nil。  

例子
```go
package main

import (
	"encoding/json"
	"fmt"
)

type greeting struct {
	SomeMessage string
}

func main() {
	// 包含原始資料的結構
	var v greeting
	v.SomeMessage = "Marshal me!"

	// 編碼成 JSON 格式字串
	json, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", json)
}
```
顯示結果：
```
{"SomeMessage":"Marshal me!"}
```
> 程式中的資料放在一個結構中，而它只有一個可匯出欄位叫 SomeMessage。  

### Marshal() 在解析結構時，會遵循以下規則來產生 JSON 鍵與值對：
* 只有可匯出的欄位（大寫字母開頭）才能被加入為 JSON 鍵。
* 帶有 JSON 標籤的欄位才會被加入，其他則忽略。
* 如果結構只有一個欄位，那不管有沒有 JSON 標籤都會加入。
* 如果結構有多重欄位，但都沒有 JSON 標籤，那麼會全數忽略（且不產生錯誤）。
