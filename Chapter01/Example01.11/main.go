//1-4 算符（operators)
//1-4-1 算符基礎
//算符是用來處理軟體資料的工具。
/*算符分類：
1.算術算符（arithmetic operators）：如四則運算。
2.比較算符（comparison operators）：用來比較兩值，如是否相等、或大小等。
3.邏輯算符（logical operators）：搭配布林值使用。
4.定址算符（address operators）：專門用來處理指標。
5.位元算符（bitwise operators）：用到的機會不多。
6.受理算符（receive operators）：用來對GO特有的通道（channe）寫入或讀取值。（lesson 16）
*/

// 用算符處理數字
package main

import "fmt"

// 例子、餐廳用餐，算服務費、來店次數
func main() {
	var total float64 = 2 * 13
	fmt.Println("+主餐：", total)

	total = total + (4 * 2.25)
	fmt.Println("+飲料：", total)

	total = total - 5
	fmt.Println("+折扣：", total)

	tip := total * 0.1
	total = total + tip
	fmt.Println("+小費：", total)

	split := total / 2
	fmt.Println("分攤額：", split)

	visitCount := 24
	visitCount = visitCount + 1

	customer := customerName()
	fmt.Println(customer)
	//計算來店次數
	remainder := visitCount % 5
	if remainder == 0 {
		fmt.Println("您已獲得來店滿五次折價券")
	}
}

// 字串連接
func customerName() string {
	givenName := "Jose"
	familyName := "John"
	return givenName + " " + familyName
}
