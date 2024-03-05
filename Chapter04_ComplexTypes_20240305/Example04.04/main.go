// 4-3-3  透過索引鍵賦值
/*
Go 語言很有彈性，允許用任何順序寫出索引鍵。
當你的陣列索引具有特殊意義，或只想對某個元素賦值，但不想動到其他元素時，這種方式就很實用。

＊＊[<長度>]<型別>{<索引鍵1>:<值1>,<索引鍵2>:<值2>...<索引鍵N>:<值N>}

*/

// 練習、以索引鍵賦予陣列初始值
package main

import "fmt"

var (
	arr1 [10]int
	arr2 = [...]int{9: 0}
	arr3 = [10]int{1, 9: 10, 4: 5} //排列順序可以跳著排
)

func comArrays() (bool, bool) {
	return arr1 == arr2, arr1 == arr3
}

func main() {
	comp1, comp2 := comArrays()
	fmt.Println("[10]int == [...]int{9:0} :", comp1)
	fmt.Println("arr2     :", arr2)
	fmt.Println("[5]int == [...]int{1,9:10,4:5} :", comp2)
	fmt.Println("arr3     :", arr3) //會是 false 是因為元素內容不一樣。
}

//如果在賦予初始值時，沒有索引鍵的值寫在其他索引鍵的中間或後面？那就是順著下去。
//例、[7]int{6:7,2:3,1} 會得到 [0,0,3,1,0,0,7,0]
