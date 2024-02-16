// 1-4-3 值的比較
/*
比較算符：==、!=、<、<=、>、>=
邏輯算符：&&、||(左右側任一為真,即真)、!(只處理單一布林值，若值為真，則為偽)
*/

package main

import "fmt"

func main() {
	visits := 15

	fmt.Println("新顧客  ：", visits == 1)
	fmt.Println("熟客    ：", visits != 1)
	fmt.Println("銀牌會員：", visits > 10 && visits <= 20)
	fmt.Println("金牌會員：", visits > 20 && visits <= 30)
	fmt.Println("白金VIP ：", visits > 30)
}
