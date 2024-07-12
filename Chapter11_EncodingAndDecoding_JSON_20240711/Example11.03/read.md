## 11-2-2 加上結構 JSON 標籤
可以給結構欄位加上標籤（tag），好讓 Unmarshal() 知道欄位要怎麼用在 JSON 解碼。  

標籤必須用原始字串（用``括住）寫在欄位後面：
```go
type person struct{
    LastName string `json:"lname"`
}
```
> 這個標籤 json 的值為 "lname"，指 LastName 欄位要對應到 JSON 資料裡的 lname 鍵。  
> 有了這個標籤，結構欄位的名稱就可以隨意命名了，但名稱需要大寫才能匯出。  

* Unmarshal() 會根據以下規則來決定要把 JSON 的鍵配對到哪個結構欄位：
    * 某個可匯出欄位的「標籤值」可以對應到 JSON 鍵。
    * 某個可匯出「欄位本身的名稱」有對應到 JSON 鍵（大小寫可不同）。 
    * 找不到符合的欄位，該 JSON 鍵就會被略過（值不會放進結構的任何欄位）。  

例子、
```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type greeting struct {
	SomeMessage string `json:"message"`
}

func main() {
	// JSON 資料
	data := []byte(`
	{
		"message":"Greetings fellow gopher!"
	}
	`)

	// 檢查 JSON 格式是否不正確
	if !json.Valid(data) {
		fmt.Printf("JSON 格式無效：%s", data)
		os.Exit(1)
	}

	v := greeting{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
```
顯示結果：
```
{Greetings fellow gopher!}
```
