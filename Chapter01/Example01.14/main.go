//1-5 零值（zero values）

//零值指的是該型別具有的預設值或空值。

/*
零值：
【型別】 				【零值】

		bool					false
		數字（整數、浮點數）	0
		string				""
	    指標、函式、介面、切片、通道、印射表	nil
*/
package main

import (
	"fmt"
	"time"
)

// 例子、宣告變數但沒給初始值，並用 Printf() 印出零值
func main() {
	var count int
	fmt.Printf("Count: %#v\n", count)
	fmt.Printf("Count: %d\n", count)

	var discount float64
	fmt.Printf("Discount: %#v\n", discount)

	var debug bool
	fmt.Printf("Debug: %#v\n", debug)
	fmt.Printf("Debug: %T\n", debug)

	var message string
	fmt.Printf("Message: %#v\n", message)

	var emails []string //字串切片
	fmt.Printf("Emails: %#v\n", emails)

	var startTime time.Time
	fmt.Printf("Start: %#v\n", startTime)
}

//Printf() 使用一種格式化樣板語言（template language），藉以轉換我們遞給他的值。
// 格式話符號 %#v ：當想以某種方式顯示變數的值或型別時用。
// 其他格式化符號：
/*
%v   任何值。若不在意印出來的值的型別時用。
%+v  印出值並加上額外資訊，例如結構型別的欄位名稱。
%#v	 用 Go 語法印出值，等於 %+v 加上型別名稱
%T	 印出值的型別
%t	 印出布林值（true/false）
%d	 印出十進位數字
%s	 印出字串
%%	 印出百分比符號
*/
