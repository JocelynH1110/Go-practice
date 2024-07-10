package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	duration1 := time.Duration(time.Second * 360)            // 時間長度 1（360 秒，等於 6 分鐘）
	duration2 := time.Duration(time.Hour*1 + time.Minute*30) // 時間長度 2（1 小時又 30 分鐘）
	fmt.Println("Dur1:", duration1.Nanoseconds(), "ns")      // 顯示時間長度值（以奈秒為單位）
	fmt.Println("Dur2:", duration2.Nanoseconds(), "ns")

	// 取得加上時間長度後的新時間
	date1 := now.Add(duration1)
	date2 := now.Add(duration2)

	// 顯示時間
	fmt.Println("Now:", now)
	fmt.Println("Date1:", date1)
	fmt.Println("Date2:", date2)
}
