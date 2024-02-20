// 1-6-4	 採用指標的函式設計
/*
若變數是指標，貨船遞給函是的是指標變數，那在函式中對該參數的值做的任何變動，也會連帶影響到函式外部原始變數的值。
*/
package main

import "fmt"

func adds5Values(count int) {
	count += 5
	fmt.Println("add5Values：", count)
}

func add5points(count *int) {
	*count += 5
	fmt.Println("add5Points：", *count)
}
func main() {
	var count int
	adds5Values(count)
	fmt.Println("add5Values post：", count)

	add5points(&count)
	fmt.Println("add5Points post：", count)
}

//以直傳遞變數時，在函式內對變數做的變動只在函式內有效，不會影響傳遞給函式的原始變數。
//以指標形式傳入函式，就真的會改變原始變數。
