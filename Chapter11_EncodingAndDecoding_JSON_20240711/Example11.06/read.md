## 11-3-2 將有多重欄位的結構轉為 JSON
如果結構非為單一欄位要轉換為 JSON 格式，需要在欄位後面加上 JSON 標籤才會被轉成鍵。  

* 例子、有些欄位未被覆值，會如何產生 JSON 字串呢？
```go
package main

import (
	"encoding/json"
	"fmt"
)

type book struct {
	ISBN          string `json:"isbn"`
	Title         string `json:"title"`
	YearPublished int    `json:"yearpub"`
	Author        string `json:"author"`
	CoAuthor      string `json:"coauthor"`
}

func main() {
	b := book{}
	b.ISBN = "9933HIST"
	b.Title = "Herry Potter"
	b.Author = "J.K"
	// 沒有對 YearPublished 和 CoAuthor

	json, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", json)
	fmt.Println(string(json))
}
```
顯示結果：
```
{"isbn":"9933HIST","title":"Herry Potter","yearpub":0,"author":"J.K","coauthor":""}
{"isbn":"9933HIST","title":"Herry Potter","yearpub":0,"author":"J.K","coauthor":""}
```
> 未被賦值的欄位仍被轉成鍵放進 JSON 資料，其值也維持 Go 語言的零值
