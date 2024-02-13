// 1-2-6 短宣告變數_一次宣告多個變數
// 用逗號隔開，:=左右數量和順序要對齊。
package main

import (
	"fmt"
	"time"
)

func main() {
	Debug, LogLevel, startUpTime := false, "info", time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
}
