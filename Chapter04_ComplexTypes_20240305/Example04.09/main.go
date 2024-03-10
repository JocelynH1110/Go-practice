// 4-4  切片（slice）
/*
陣列:
很方便，但它對長度的硬性規定會帶來一些麻煩。
如要寫一個函式來接收一個陣列和處理其資料，那麼這個函式就只能處理特定長度的陣列。
萬一有陣列長度不同，就得寫新的函式來因應。
陣列在處理有序資料即得時很方便。

切片:
是在陣列外頭套上一層額外包裝，能和陣列一樣建立有數字索引鍵的有序集合，卻不必擔心長度問題。（Go 語言會處理好細節，例、動態調整陣列長度）
另一個優勢，可以用 Go 語言內建函式 append() 新增切片元素。

＊＊這函是的輸入值是你的切片和你要加入的值，它會輸出添加了新元素的新切片:
<新切片> = append(<切片>, <新元素>)

很多時候會先宣告一個空切片，然後取得資料後再慢慢擴充它。
事實上大部分程式碼都鮮少使用陣列。
只有當得把元素限制在某個數量時，才會用到陣列。
切片更容易拿來在函式之間傳遞。

陣列和切片的共通處：
1.同樣只能容納單一型別的元素。
2.可用[]去讀寫任一元素。
3.也能用 for i 迴圈走訪它。
*/

// 4-4-1  使用切片

// 練習、建立與使用切片
package main

import (
	"fmt"
	"os"
)

func getPassedArg(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Printf("至少需要輸入 %v 個參數\n", minArgs)
		os.Exit(1) //強制結束程式
	}
	var args []string //建立空白切片

	for i := 1; i < len(os.Args); i++ { //因為要去掉第一個參數（程式名稱），所以迴圈從 1 開始
		args = append(args, os.Args[i])
	}
	return args
}
func findLongest(args []string) string {
	var longest string
	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}
	return longest
}
func main() {
	if longest := findLongest(getPassedArg(4)); len(longest) > 0 {
		fmt.Println("傳入的最長單字", longest)
	} else {
		fmt.Println("發生錯誤")
		os.Exit(1)
	}

}
