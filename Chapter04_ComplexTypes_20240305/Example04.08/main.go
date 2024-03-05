// 4-3-6  走訪一個陣列

// 練習、用 for i 迴圈走訪和處理陣列--參數版
// 在這練習中要把陣列傳給函式，而函式會對陣列做些處理後傳回。為了能處理相同的陣列，函式的參數和傳回值也必須指定同樣的陣列長度。
package main

import "fmt"

func fillArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}

func opArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] + arr[i]
	}
	return arr

}
func main() {
	var arr [10]int
	arr = fillArray(arr)
	arr = opArray(arr)
	fmt.Println(arr)
}
