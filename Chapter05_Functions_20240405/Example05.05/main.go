// 5-3  參數不定函式（variadic function）
/*
指的是一個函式可以接收數量不確定的參數。
若無法確定某參數要接收的引數有多少個，就是使用參數不定函式的時機。

＊＊func f(參數名稱 ...型別)

打包算符（pack operator）：型別前的三個點 ... 。
此算符讓函式參數變成可接收數量不定的引數。
他的作用在於告訴 Go 語言，把所有符合該型別的引數都放進此參數名稱，打包成一個切片。
它可接收任意數量的引數，甚至是完全沒有引數。

例子、讓函式接收數量不定的引數
package main

import "fmt"

func main() {
	nums(99, 100)
	nums(200)
	nums()
}

func nums(i ...int) {
	fmt.Println(i)
}

//參數不定函式規則：
1.數量不定的參數必須放在所有數量固定的參數的最後面。(如果不定參數在前面的話，Go 語言就無法分辨哪些引數是要傳給後面的參數)
2.一個函式只能容許一組數量可變得參數。

＊＊不固定參數放最後面的寫法：
func main(){
	nums("Fafa",20,100)
}

func nums(person string,i ...int){
	fmt.Println(person)
	fmt.Println(i)
}
*/

/*
將切片元素傳給參數不定函式：
會產生錯誤，因為函式等待接收的是一連串型別為 int 的引數，然後再把它們收集成一段 []int 切片。
有一種解法：
解包算符（unpack operator），將切片引數的元素拆解出來、一一傳給數量可變參數，就是在切片變數後面使用... 。

NOTE、只有切片可以用解包算符。長度固定的陣列會被認為其型別與參數型別不同。但可以用「陣列[:]...」的語法把它轉成切片傳入數量不定參數。

例子、將切片元素拆解傳遞給不定參數
package main

import "fmt"

func main() {
	i := []int{2, 5, 10}
	nums(i...) //相當於呼叫nums(2,5,10)
}
func nums(i ...int) {
	fmt.Println(i)
}
*/

// 練習、數值加總函式
package main

import "fmt"

func main() {
	i := []int{5, 10, 15}
	fmt.Println(sum(5, 4))
	fmt.Println(sum(i...))
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
