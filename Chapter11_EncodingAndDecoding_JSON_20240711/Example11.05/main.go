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
