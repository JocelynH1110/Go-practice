package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct { // 父結構
	Lastname  string  `json:"lname"`
	Firstname string  `json:"fname"`
	Address   address `json:"address"` // 子結構型別欄位
}

type address struct { // 子結構
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

func main() {
	// JSON 資料
	data := []byte(`
	{
	"lname":"Smith",
	"fname":"John",
	"address":{
		"street":"Daan Road",
		"city":"Taipei",
		"state":"Taiwan",
		"zipcode":106
		}
	}
	`)
	dataStr := string(data)
	// 解析 JSON 並將值存入結構 p
	p := person{}

	// 用 strings.NewReader() 從字串建立一個 io.Reader
	// 並以此建立 json,Decoder
	decoder := json.NewDecoder(strings.NewReader(dataStr))
	// Decoder 將 JSON 編碼和轉換成 結構p
	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(p)
	fmt.Println()

	/* 使用 Unmarshal()
	if err := json.Unmarshal(data, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", p) // 這樣可以印出欄位
	*/

	// 建立 json.Encoder，寫入對象是 os.Stdout（主控台）
	encoder := json.NewEncoder(os.Stdout)
	// 設定前綴詞和縮排字串
	encoder.SetIndent(",", "\t")
	// 將結構 p 編碼為 JSON
	if err := encoder.Encode(p); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
