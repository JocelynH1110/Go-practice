// 1-2-2.用 var 一次宣告多個變數 => var (名稱1 型別1 =值1  名稱2 型別2 =值2...)

// 這種宣告方式在宣告套件範圍的變數時很常見。變數型別不必一樣，還可以有各自初始值。
// 函式裡也可以這樣宣告，但比較少人這樣做。
package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool      = false
	LogLevel    string    = "info"
	startUpTime time.Time = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
