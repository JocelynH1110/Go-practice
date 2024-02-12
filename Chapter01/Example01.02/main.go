// 1-2 宣告變數
// 變數可用來暫存資料。
// 宣告變數需要滿足四個條件：
//
//	1.宣告變數的敘述
//	2.變數名稱
//	3.變數要儲存的資料型別
//	4.變數的初始值

// 1-2-1.用 var 宣告變數 => var 變數名稱 變數型別 = 值
package main

import (
	"fmt"
)

var foo string = "bar"

func main() {
	var baz string = "qux"
	fmt.Println(foo, baz)
	//Println()可以傳入多個變數或值，並用逗號分開，他們印出來時各值之間會隔一個空白。
}

//補充：有宣告但未使用的變數在編譯時會產生『declared but not used』錯誤。這是為了確保開發者不要宣告一堆沒有用的變數，浪費記憶體空間。
