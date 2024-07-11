package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second * 2)
	end := time.Now()
	duration1 := end.Sub(start)    // 計算兩個時間值之間的長度
	duration2 := time.Since(start) // 計算 start 到 time.Now() 的時間長度

	fmt.Println("Duration1:", duration1)
	fmt.Println("Duration2:", duration2)

	// 檢查 duration1 是否小於 2500 毫秒（2.5秒）
	if duration1 < time.Duration(time.Millisecond*2500) {
		fmt.Println("程式執行時間符合預期")
	} else {
		fmt.Println("程式執行時間超出預期")
	}
}
