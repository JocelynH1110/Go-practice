package main

import (
	"fmt"
	"io"
	"os"
)

// 檢查檔案是否存在的自訂函式
func main() {
	f, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("檔案內容：")
	fmt.Println(string(content))
}
