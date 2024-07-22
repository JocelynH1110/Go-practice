package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	v := make(map[string]interface{}) // 初始化 map

	// 存入原始資料
	v["checkNum"] = 123
	v["amount"] = 300
	v["category"] = []string{"gift", "clothing"}

	// 將 map 編碼成 JSON 格式
	jsonData, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))
}
