// 練習、載入預算分類
// 在執行 main() 函式之前將一系列的支出預算分類存入一個 map 集合，以便讓 main() 只須負責印出 map 內容。

// 此處目的在展示 init() 如何能用於資料初始化，讓 main() 內容更加簡潔，且有現成資料。
package main

import (
	"fmt"
)

var budgetCategories = make(map[int]string)

func init() {
	fmt.Println("初始化 budgetCategories...")
	budgetCategories[1] = "汽車保險"
	budgetCategories[2] = "房屋貸款"
	budgetCategories[3] = "電費"
	budgetCategories[4] = "退休金"
	budgetCategories[5] = "雜貨支出"
	budgetCategories[7] = "旅遊補助"
	budgetCategories[8] = "汽車貸款"
}

func main() {
	for k, v := range budgetCategories {
		fmt.Printf("鍵：%d, 值：%s\n", k, v)
	}
}
