// 練習：載入預算分類
package main

import "fmt"

var budgetGategories = make(map[int]string)

func init() {
	fmt.Println("初始化 budgetGategories...")
	budgetGategories[1] = "汽車保險"
	budgetGategories[2] = "房屋貸款"
	budgetGategories[3] = "汽車保險"
	budgetGategories[4] = "電費"
	budgetGategories[5] = "汽車貸款"
	budgetGategories[6] = "旅遊補助"
	budgetGategories[7] = "雜貨支出"
}

func main() {
	for k, v := range budgetGategories {
		fmt.Printf("鍵： %d,值：%s\n", k, v)
	}
}
