// 4-3-6  走訪一個陣列
/*
最常操作陣列的方式會是透過迴圈。
由於索引必然是從 0 連續遞增到陣列長度減 1，用迴圈來走訪（iterate，也稱迭代）陣列就非常容易。

使用 len() 取得陣列長度的效率：
陣列長度是陣列型別定義的一部分。Go 會自動追蹤其的元素數量，大大提高使用 len() 的效率，切片和 map 也是。
*/

// 練習、用 for i 迴圈走訪和處理陣列
// 定義一個陣列，並用若干數字賦予初始值。用迴圈一一走訪和處理這些值，並把結果放進一個訊息。這訊息會回傳並印出來。
package main

import "fmt"

func message() string {
	m := ""
	arr := [4]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
		m += fmt.Sprintf("%v: %v\n", i, arr[i]) //將用格式化字串將索引和值一一連起來
	}
	return m
}

func main() {
	fmt.Print(message())
}
