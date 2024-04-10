// 5-4-2 建立閉包（closure）
/*
閉包：
是匿名函式的諸多形式之一。
一般函式在離開某個函式範圍後，就沒辦法繼續引用父函式的區域變數，可是閉包卻能。

以下為看似正常的匿名函式：
package main

import "fmt"

func main() {
	i := 0
	increment := func() int {
		i++
		return i
	}

	fmt.Println(increment())
	fmt.Println(increment())
	i += 10
	fmt.Println(increment())
}
//以上匿名函式 increment() 會把父函式 main() 的變數 i 遞增 1 並傳回，在每次 main() 呼叫它時，都可以看到 i 的值改變。


package main

import "fmt"

func main() {
	increment := incrementor()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//當從 main() 呼叫 incrementor() 傳回的匿名函式 increment() 時，會發現它居然記得父函式 incrementor() 的區域變數 i，儘管 incrementor() 已經執行完畢了。
// increment 就是所謂的語意閉包（lexical closure）或簡稱閉包，因為這函式「包住了」它所引用的外部變數。換言之、閉包能記住父函式的變數，即使離開了父函式的執行範圍也一樣。
*/

// 練習、建立一個閉包函式來製作倒數計數器
package main

import "fmt"

func main() {
	max := 4

	counter := decrement(max) //取得閉包函式
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func decrement(i int) func() int {
	return func() int { //閉包會記住父函式的參數 i
		if i > 0 { //若 i 仍大於 0 就遞減
			i--
		}
		return i
	}
}
