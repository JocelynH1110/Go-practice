//  3-4 字串（String）
// 3-4-1 字串與字串常值（string literal）

//Go 語言只有一種文字型別，就是 String 。
/*
其支援兩種字串常值：
1.原始的（raw）：由一對反引號``括住的字串。
2.轉譯的（interpreted）：由一對雙引號""括住的字串。

若你的字串變數儲存的是原始字串時，變數內容就會跟字串在螢幕上的內容完全一樣。唯獨無法顯示反引號``。
但若是轉譯字串時，Go 語言會先掃過你寫的內容，並用他的規則轉換某些文字。例如有\n 之類的會換行。
*/

// 例、兩種字串顯示效果
package main

import "fmt"

func main() {
	comment1 := `This is the BEST 
	thing ever!`
	comment2 := `This is the BEST\nthing ever!`
	comment3 := "This is the BEST\nthing ever!"

	fmt.Print(comment1, "\n\n")
	fmt.Print(comment2, "\n\n")
	fmt.Print(comment3, "\n")

	//如果你想顯示的文字有大量換行字元、雙引號、反斜線字元，那用原始字串會較方便。
	comment4 := `In "Windows" the user directory is "C:\Users\"`
	comment5 := "In \"Windows\" the user directory is \"C:Users\\\"" //在轉譯字串中要正確顯示雙引號和反斜線，前面必須多寫一次反斜線。

	fmt.Println(comment4)
	fmt.Println(comment5)
}

//字串常值只是用來把文字存進 string 型別變數的辦法。一存進去變數後，不管用什麼辦法存的就都沒啥差別了。
