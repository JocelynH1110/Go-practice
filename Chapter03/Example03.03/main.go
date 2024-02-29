// 3-3-2   浮點數 → float32、float64
// float64 容量較大，精準度也較高。

/*
浮點數在儲存數值時，會將數值拆成整數和小數部份。
整數和小數各佔多少位元，要看數值本身而定。
例、9999.9，儲存整數所需的位元就會比小數的要多
*/

// 練習：浮點數的準確度
package main

import "fmt"

func main() {
	var a int = 100
	var b float32 = 100
	var c float64 = 100

	fmt.Println(a / 3)
	fmt.Println(b / 3)
	fmt.Println(c / 3)
}

//以上練習可看出，電腦無法對除不盡的運算給出完美答案，但 float64 比 float32 要精確更多。
//為何結尾小數點會是 3 以外數字？因為電腦使用二進位來儲存小數位，因此多少一定會有些誤差。
