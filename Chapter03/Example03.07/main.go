// 3-4-2  Rune
// rune（符文）是一種具備充足空間、足以容納單一一個 UTF-8 字元的型別。（Unicode 編碼，會佔用 1-4 個位元組不等）
// 在 Go 語言中，字串常值都是用 UTF-8 來編碼。

/*
舊式標準如 ASCII 只用一個位元組來編碼，UTF-8 最多會用到 4 個位元組。
當文字以 string 型別儲存時，Go 語言會以 byte 集合來儲存所有字串（string 實際上便是唯讀的 byte 切片）。
這意味著有些 UTF-8 字元會被拆成多個位元組。

＊＊為了能安全處理任何字串，不論其編碼方式是採用單一或是多重位元組，最好是把字串從 byte 集合轉換成 rune 集合。

會用到多重位元組的字元，都是諸如：中文字、日文字、特殊拉丁語系字母等。
*/

// 例、用 Go 來處理字串的個別位元組：
package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	for i := 0; i < len(username); i++ {
		fmt.Print(username[i], " ")
	}
	fmt.Print("\n")

	//將位元組轉回字串
	for y := 0; y < len(username); y++ {
		fmt.Print(string(username[y]), " ")
	}
	fmt.Print("\n")
}

//本來定義的字串是 13 個字元，但因其夾雜了一個由雙位元組編碼的字元 Ü，導致印出來有 14 個數值。
//用函式將每個位元組轉回字元時 Ü 的兩個位元編碼被拆開解讀，所以就會出錯。

//＊＊為了能安全的處理多位元組編碼字串的每一個字元，必須先把 byte 型別的字串切片轉換成 rune 型別的切片。
