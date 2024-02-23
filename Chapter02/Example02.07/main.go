package main

import "fmt"

//練習、使用 for i 迴圈
/*
func main() {
	//在迴圈建立變數i，初始值為0 ,在 i 小於 5 時繼續重複迴圈，每次迴圈結束後遞增1
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
*/

// 以上例子寫死了結束的條件（i<5）；然而真正在走訪陣列和切片時，此條件往往由集合的元素數量決定。
// 練習、用 for 迴圈走訪切片元素
func main() {
	names := []string{"Dora", "Mika", "Cece"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
