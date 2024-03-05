// 3-4-2 rune
//檢查字串長度

package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	fmt.Println("Bytes", len(username))
	fmt.Println("Runes", len([]rune(username)))

	//用切片擷取字串前 10 個元素，剛好到多位元組字元
	fmt.Println(string(username[:10]))
	fmt.Println(string([]rune(username)[:10]))
}

//如果直接對（含有多位元組字元的）字串施以 len() 函式處理，會得到錯誤的數目。
//故當你要處理 string 變數，且需要計算長度或擷取特定數量的字元時，應該先轉換為 rune 切片。
