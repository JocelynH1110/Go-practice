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
