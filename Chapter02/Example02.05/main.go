// 2-3-2 switch 的不同用法
// case 後面的值或運算式，愛寫幾個都可以，只要用逗號分開即可；Go 語言會由左到右檢查這些值或運算式。

// 例、 switch 敘述和多重 case 配對值
package main

import (
	"fmt"
	"time"
)

/*
func main() {
	dayBorn := time.Sunday
	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("生日為平日")
	case time.Saturday, time.Sunday:
		fmt.Println("生日為週末")
	default:
		fmt.Println("生日錯誤")
	}
}
*/

// 例、沒有運算式的 switch 敘述
func main() {
	switch dayBorn := time.Sunday; { //此為起始賦值敘述
	case dayBorn == time.Sunday || dayBorn == time.Saturday:
		fmt.Println("生日為週末")
	default:
		fmt.Println("生日為非週末")
	}
}
