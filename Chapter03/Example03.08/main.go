// 3-4-2 rune
// 將字串切片轉換成 rune 切片（為了安全走訪每個字）

package main

import "fmt"

func main() {
	username := "Sir_King_Über"
	runes := []rune(username)

	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}
}
