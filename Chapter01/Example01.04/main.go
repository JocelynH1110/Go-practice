// 1-2-3 省略型別或值的宣告變數
package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		Debug       bool
		LogLevel    = "info"
		startUpTime = time.Now()
	)

	fmt.Println(Debug, LogLevel, startUpTime)
}
