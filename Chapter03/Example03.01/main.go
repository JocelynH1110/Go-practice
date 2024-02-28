// 3-1  核心型別--前言
//布林值、數字（整數、浮點數）、字串
/*
Go 語言是強型別（strongly typed）語言，意即所有資料都必須屬於某個型別，這型別必須要是固定的、無法改變的。
變數可以隨意轉換型別，又被稱作弱型別（weakly typed）語言。

另一種分法：
靜態型別（statically typed）：會在編譯時檢查型別。（golang）
動態型別（dynamically typed）：在執行階段檢查型別。

程式語言必須明確知道一個值是文字或數字，才好確定它能用到多少記憶體空間。
*/

// 3-2  布林值（bool）：true/false
/*
bool的值只有 true 和 false 兩種，其零值為 false 。
當使用 == 、> 這類算符，結果一定是 bool 值。
*/

package main

import "fmt"

func main() {
	fmt.Println(10 > 5)
	fmt.Println(10 == 5)
}
