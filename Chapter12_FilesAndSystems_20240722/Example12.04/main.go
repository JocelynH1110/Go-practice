package main

import "os"

func main() {
	f, err := os.Create("text.txt") // 建立文字檔
	if err != nil {
		panic(err)
	}
	defer f.Close() // 確保在 main() 結束時關閉檔案
	f.Write([]byte("使用 Write() 寫入\n"))
	f.WriteString("使用 WriteString() 寫入\n")
}
