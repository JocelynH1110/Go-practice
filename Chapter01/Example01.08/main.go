// 用 var 宣告多種變數
package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	//型別
	var start, middle, end float32
	fmt.Println(start, middle, end)

	//不同型別的值
	var name, left, right, top, bottom = "one", 1, 1.5, 2, 2.5
	fmt.Println(name, left, right, top, bottom)

	//函式
	var Debug, LogLevel, startUpTime = getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}
