package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() // 取得當下系統時間
	fmt.Println("程式開始時間：", start)
	fmt.Println("資料處理中...")
	time.Sleep(2 * time.Second) // 讓程式停頓 2 秒，模擬資料處理
	end := time.Now()           // 再次取得當下系統時間
	fmt.Println("程式結束時間：", end)
}
