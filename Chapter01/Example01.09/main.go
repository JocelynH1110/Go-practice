// 1-3 更改變數值
package main

import "fmt"

//改變單一變數的值
/*
func main() {
	offset := 5
	fmt.Println(offset)

	offset = 10
	fmt.Println(offset)
}
*/

// 用其他變數來賦值
var defaultOffset = 10

func main() {
	offset := defaultOffset
	fmt.Println(offset)

	offset = offset + defaultOffset
	fmt.Println(offset)
}
