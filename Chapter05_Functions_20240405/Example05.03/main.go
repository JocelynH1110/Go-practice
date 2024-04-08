// 5-2-3 函式回傳值
/*
＊＊接收多重傳回值
值1, 值2...:=函式名稱()

函式通常會接收輸入值、做若干處理後傳回處理結果。
有傳回值的函式，裡面必須有 return 敘述。
*/

package main

import "fmt"

func fizzBuzz(i int) (int, string) {
	switch {
	case i%15 == 0:
		return i, "FizzBuzz"
	case i%3 == 0:
		return i, "Fizz"
	case i%5 == 0:
		return i, "Buzz"
	}
	return i, ""
}

func main() {
	for i := 1; i <= 15; i++ {
		n, s := fizzBuzz(i)
		fmt.Printf("Results: %d %s\n", n, s)
	}
}

//忽略一部分傳回值
/*
假設我們不想收到上面函式（fizzBuzz()）的兩個值，但又不能只接收單一個值，這時就可以用一個空白符號（blank identifier，即底線）來忽略該傳回值。

例如、_,s :=fixxBuzz(i)

NOTE、用 := 短變數宣告賦值時，左邊一定至少要有一個實際變數。不然編譯器就會產生錯誤。
*/
