// 練習、將收款方與預算分類做配對

package main

import (
	"fmt"
)

var budgetCategories = make(map[int]string)
var payeeToCategory = make(map[string]int)

func init() { //第一個執行的 init()
	fmt.Println("初始化 budgetCategories...")
	budgetCategories[1] = "汽車保險"
	budgetCategories[2] = "房屋貸款"
	budgetCategories[3] = "電費"
	budgetCategories[4] = "退休金"
	budgetCategories[5] = "雜貨支出"
	budgetCategories[7] = "旅遊補助"
	budgetCategories[8] = "汽車貸款"
}

func init() { // 第二個執行的 init()
	fmt.Println("設定收款人與其預算分類...")
	payeeToCategory["Nationwide"] = 1
	payeeToCategory["BBT Loan"] = 2
	payeeToCategory["First Energy Electric"] = 3
	payeeToCategory["Ameriprise Financial"] = 4
	payeeToCategory["Walt Disney World:]"] = 5
	payeeToCategory["ALDI"] = 7
	payeeToCategory["Martins"] = 7
	payeeToCategory["Wal Mart"] = 7
	payeeToCategory["Chevy Loan"] = 8
}

func main() {
	fmt.Println("主程式：印出收款人與預算分類名稱")
	for k, v := range payeeToCategory {
		fmt.Printf("收款人：%s, 分類：%s\n", k, budgetCategories[v])
	}
}
