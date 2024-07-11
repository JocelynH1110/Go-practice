# 11-2 解碼 JSON 為 Go 結構
解碼JSON：指將 JSON 資料轉換成 Go 的資料型別。  
Go 語言會自行將 JSON 值轉成對應的 Go 語言型別，讓我們得以用 Go 語言的方式處理資料。  

事先知道 JSON 包含哪些鍵（key），就可以解析（unmarshal）它、再把結果存在一個相對應的 Go 結構中，這得用到 Go 標準函式庫 encoding/json 的 Unmarshal() 函式。

## 11-2-1 Unmarshal()
```
func Unmarshal(data []byte, v interface{}) error
```
> * 參數 data([]byte 切片)：儲存 JSON 資料的字串。
> * v ：用來儲存解析結果的變數，其型別為空介面，會傳入一個結構指標。
> * Unmarshal() 會解析 JSON 字串，並試著將結果存到該結構中。
> * v 不能為 nil，否則會傳回「call of Unmarshal passes non-pointer as second argument」（呼叫 Unmarshal 時傳入非指標）的錯誤。  

* 如何將 JSON 資料轉成結構：
```go
package main

import (
	"encoding/json"
	"fmt"
)

type greeting struct { // 用來儲存解碼的 JSON 資料結構
	Message string
}

func main() {
	// JSON 資料
	data := []byte(`
	{
		"message":"Greetings fellow gopher!"
	}
	`)

	var v greeting                  // 建一個空結構 v
	err := json.Unmarshal(data, &v) // 解析 JSON 和寫入 v
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)         // 印出 v 的內容
	fmt.Printf("%#v\n", v) // 可看出結構名稱和欄位名稱
}
```
結果顯示：
```
{Greetings fellow gopher!}
main.greeting{Message:"Greetings fellow gopher!"}
```

p.s.結構欄位必須是可匯出的（exportable），也就是名稱首字用英文大寫，才能夠被 Unmarshal() 看到和使用。未匯出（首字小寫）的欄位會被忽略。
