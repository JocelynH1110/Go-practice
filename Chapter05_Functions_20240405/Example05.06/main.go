// 5-4 匿名函式與閉包
/*
具名函式（named functions）：自帶識別名稱，且需在套件層級宣告。
匿名函式（anonymous functions）：沒有名稱且只能在其他函式內宣告，也只能使用一次（除非在建立它後指派給一個變數，才能重複呼叫）。
以上兩者主要差別在於宣告時不會寫函式名稱，其他連接引數、傳回值都相同。

匿名函式可搭配以下目的或功能：
1.定義只使用一次的函式
2.定義要傳回給另一個函式的函式
3.定義 Goroutine 的程式碼區塊（lesson 16）
4.實作閉包
5.搭配 defer 敘述延後執行程式碼
*/

// 5-4-1  宣告匿名函式
/*
//// 宣告匿名函式＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝
func main() {
	func() {
		fmt.Println("Greeting")
	}() //用() 立即呼叫它
}

//執行小括號（execution parenttheses）：在函式又大括號後面，會當場呼叫匿名函式並執行它。


//// 要傳給函式的引數必須寫在執行小括號＝＝＝＝＝＝＝＝＝＝＝＝
func main() {
	message := "greeting"

	func(str string) {
		fmt.Println(str)
	}(message)
}


//// 將匿名函式傳給變數，已利再次使用：＝＝＝＝＝＝＝＝＝＝＝＝
func main() {
	f := func() {
		fmt.Println("透過變數呼叫一個匿名函式")
	}
	fmt.Println("匿名函式宣告的下一行")
	f() //透過變數 f 呼叫匿名函式
}
*/

// 練習、建立一個匿名函式來計算數值平方
package main

import "fmt"

func main() {
	x := 9
	sqr := func(i int) int {
		return i * i
	}
	fmt.Printf("%d 的平方為 %d\n", x, sqr(x))
}

//當我們在函式中需要一個小函式、且在煮成是的其他部位可能不需要重新利用時，就可以直接建立一個匿名函式，並將其賦予給變數。
