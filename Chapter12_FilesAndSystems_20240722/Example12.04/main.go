package main

import "os"

func main() {
	msg := "Hello Golang!"
	// 建立檔案並寫入資料
	err := os.WriteFile("text.txt", []byte(msg), 0644)
	if err != nil {
		panic(err)
	}
	rm := os.Remove("test.txt")
	if rm != nil {
		panic(rm)
	}
}
