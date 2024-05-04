// 6-5  panic

// 6-5-1 何謂 panic ? ===============================================================================
/*
大多數的程式語言都會用例外的形式處理錯誤。
Go 大多情況下錯誤會以 error 值的形式傳回，這也通常不會影響程式運作。
但若遇到真正嚴重的狀況，Go 還是會引發 panic 。
和例外有些類似的是， panic 會在函式中一路往上傳、最終使整個程式當掉，除非你處理它。

Note、發生 panic 時，會在錯誤訊息中看到像是「Goroutine running」的字樣。這是因為 main() 函式自己也是一個 Goroutine。

發生 panic ，代表程式遭遇了完全不正常的狀況。
引發 panic ，通常是要保護程式的完整性（integrity），藉由中斷程式來避免造成其他影響。
*/

// 6-5-2 panic() 函式 =================================================================================
/*
panic 也可由開發者在程式執行期間觸發，辦法是使用 panic() 函式。
它接收一個空介面型別 interface{} 參數，這意味 panic() 能接收任何型別的資料。（lesson 4）
但在大部份情況下，你應該傳入一個 error 介面型別的值，這也是 Go 的既定慣例。

當 panic 發生時，一般會伴隨以下動作：
1.停止程式執行
2.發生 panic 的函式中若有延後執行（deferred）的函式，他們會被呼叫。
3.發生 panic 的函式的上層函式中，若有延後執行（deferred）的函式，他們會被呼叫。
4.沿著函式堆疊一路往上，最後抵達 main()。
5.發生 panic 的函式之後的所有敘述都不會執行。
6.程式當掉。
*/

// 手動引發 panic
//例子、panic 的原因，因為走訪的切片索引範圍超過元素數量，Go 認為這使程式進入不正常狀況，因而觸發 panic
package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8}
	total := 0
	for i := 0; i <= 10; i++ {
		total += nums[i]
	}
	fmt.Println("總和：", total)
}


// 例子、主動使用 panic() 來讓程式當掉
package main

import (
	"errors"
	"fmt"
)

func main() {
	msg := "good bye"
	message(msg)
	fmt.Println("這行不會印出")
}

func message(msg string) {
	if msg == "good bye" {
		panic(errors.New("粗事惹"))
	}
}
