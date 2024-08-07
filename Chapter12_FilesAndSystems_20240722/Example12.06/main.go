package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

const (
	firstName = iota // CSV 欄位索引
	lastName
	age
)

func main() {
	file, err := os.Open("data.csv") // 開啟 CSV 檔案
	if err != nil {
		panic(err)
	}
	defer file.Close()

	header := true                // 標頭開關
	reader := csv.NewReader(file) // 取得 csv.Reader 結構
	for {
		record, err := reader.Read() // 從 csv.Reader 讀取一行資料
		if err == io.EOF {           // 遇到檔案結尾錯誤，就離開迴圈
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if header {
			header = false
			continue // 跳過第一行（標頭）
		}
		fmt.Println("--------------")
		fmt.Println("First name:", record[firstName])
		fmt.Println("Last name:", record[lastName])
		fmt.Println("Age:", record[age])
	}
}
