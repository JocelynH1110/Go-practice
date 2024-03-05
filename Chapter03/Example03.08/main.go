// 3-4-2 rune

// 將 byte 的字串切片轉換成 rune 切片（為了安全走訪每個字）

// 當Go 語言編譯器發現你在嘗試走訪 rune 切片時，會自動把它轉成 range 迴圈，故也可以直接用 for range 走訪 rune 。

package main

import "fmt"

func main() {
	username := "Sir_King_Über"
	runes := []rune(username)

	//用 rune 切片走訪
	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}
	fmt.Print("\n")

	//用 for range 走訪
	for _, v := range runes {
		fmt.Print(string(v), " ")
	}
}
