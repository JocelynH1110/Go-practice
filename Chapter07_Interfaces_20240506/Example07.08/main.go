// 7-4-3 空介面 interface{}
/*
空介面：就是一個沒有任方法的集合、亦即沒有定義任何行為的介面。

＊＊空介面寫法：
interface{}

空介面未指定任何方法，表示 Go 語言的任何型別都會自動實作空介面。即任何型別都能滿足空介面的規範。
*/

// 例、函式如何透過空介面，來接收任意型別的傳入值：
package main

import "fmt"

type cat struct {
	name string
}

func main() {
	i := 99
	b := false
	str := "test"
	c := cat{name: "Tida"}
	printDetails(i, b, str, c)
}

func printDetails(data ...interface{}) { //接收數量不定（...）的空介面參數
	for _, i := range data {
		fmt.Printf("%v,%T\n", i, i) //印出值和型別
	}
}

// printDetails() 函式接收一個數量不定的參數 data （會成為一個切片），其型別為空介面型別。傳入的每個值都有各自的型別，但都自動實作了 interface{} 空介面型別。

//補充、Go 語言的泛型
/*
泛型（generic）：允許函式接收不確定型別的參數，這型別等到呼叫時才會確定。
Go 語言泛型（稱為 Type Parameters）會比空介面更強大。
*/
