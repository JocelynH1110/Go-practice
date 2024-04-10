// 5-6-3  defer 對變數值的副作用
/*
使用 defer 敘述務必審慎。
其中一個必須考量到的是，若 defer 函式有使用到外部變數，它執行時會發生怎樣的結果？

當變數傳給被延後的函式時，函式會取得變數在傳遞「那一刻當下」的值。
就算變數值在該函式之後有所變動，等到 defer 函式實際執行時，它看到的變數值也不會反應外圍函式中的變動
*/
package main

import "fmt"

func main() {
	age := 25
	name := "Rara"
	defer personAge(name, age)

	age *= 2
	fmt.Println("年齡加倍：")
	personAge(name, age)
}

func personAge(name string, i int) {
	fmt.Printf("%s 是 %d 歲\n", name, i)
}
