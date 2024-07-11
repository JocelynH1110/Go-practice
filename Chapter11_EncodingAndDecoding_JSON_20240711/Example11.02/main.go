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
