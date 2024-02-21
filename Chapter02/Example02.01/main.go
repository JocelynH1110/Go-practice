//2-2	if 敘述

/* 2-2-1 if 敘述基礎
if 敘述是 Go 語言中最基本的條件判斷功能。其會根據布林運算式（Boolean expression）的回傳值決定是否執行某一區塊的程式語句。
其語法：
if <布林運算式>{
	<程式區塊>
}
若布林運算式為 true ，程式區塊便會被執行。
if 敘述只能在函式的範圍中使用。
*/

/*
2-2-2 else 敘述
當 if 布林值不成立，沒有執行第一區塊時，else 才會執行
*/
package main

import "fmt"

func main() {
	intput := 5
	if intput%2 == 0 {
		fmt.Println(intput, "是偶數")
	} else {
		fmt.Println(intput, "是奇數")
	}
}
