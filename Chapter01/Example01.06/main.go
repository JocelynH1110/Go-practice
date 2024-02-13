// 1-2-5 短宣告變數
// 可以省略 var、型別，型別可以從值推斷出來。
// 此寫法只能寫在函式裡面。
package main

import (
	"fmt"
	"time"
)

func main() {
	Debug := false
	LogLevel := "info"
	startUpTime := time.Now()

	fmt.Println(Debug, LogLevel, startUpTime)
}
