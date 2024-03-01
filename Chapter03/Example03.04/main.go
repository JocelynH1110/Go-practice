// 3-3-3   溢位和越界繞回
// 如果試著在建立變數時，賦予一個超過型別容許上限的初始值，就會發生溢位（overflow）錯誤。

/*
溢位（overflow）：在設定變數時，將其值設定超過最大值。
越界繞回（wraparound）：在建立變數後，才將其值設到超過最大值。（這編輯器無法攔截到）
*/

// 例、以下為溢位寫法
package main

import "fmt"

func main() {
	var a int8 = 128
	fmt.Println(a)
}
