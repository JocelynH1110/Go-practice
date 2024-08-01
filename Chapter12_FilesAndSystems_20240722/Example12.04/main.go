package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("檔案內容：")
	// 建立一個 bufio.Reader 結構，緩衝區大小 10
	reader := bufio.NewReaderSize(file, 10)
	for {
		// 讀取 reader 直到碰到換行符號為止（讀取一行文字）
		line, err := reader.ReadString('\n')
		fmt.Printf("%s\n", line)
		if err == io.EOF { // 若已讀到檔案結尾就結束
			break
		}
	}
}
