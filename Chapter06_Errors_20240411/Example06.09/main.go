// 6-6 recover（恢復）
// Go 語言提供了 recover() 函式，可以在某個 Goroutine 發生 panic 後取回其控制權。

/*
＊＊recover() 函式定義如下：
func recover() interface{}

recover() 函式沒有參數，傳回值是一個空介面，這意味著傳回資料可以是任意型別。
實際上，recover() 傳回的會是你一開始傳給 panic() 函式的值。

若在延遲執行的函式中呼叫 recover() ，就可以恢復正常執行和停止 panic()。
若在延遲執行的函式外呼叫 recover()，就無法阻止 panic 發生。
*/

// 例子、在使用 recover()、panic() 和 defer 時會經歷的過程：
package main

import (
	"errors"
	"fmt"
)

func main() {
	a()
	fmt.Println("這一行現在會印出了")
}

func a() {
	b("Good-night")
	fmt.Println("返回 a()")
}

func b(msg string) {
	defer func() {
		//若有 panic 發生，用 recover() 救回程式
		if r := recover(); r != nil {
			fmt.Println("b() 發生錯誤：", r)
		}
	}()
	if msg == "Good-night" {
		panic(errors.New("又禿垂了"))
	}
	fmt.Print(msg)
}

// 練習、從 panic 中復原，並調查造成 panic 的錯誤是什麼，在根據錯誤原因印出有意義的訊息給使用者看
package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("無效的時薪")
	ErrHoursWorked = errors.New("無效的一週工時")
)

func main() {
	pay := payDay(100, 25)
	fmt.Printf("週薪：%d\n\n", pay)

	pay = payDay(100, 200)
	fmt.Printf("週薪：%d\n\n", pay)

	pay = payDay(60, 25)
	fmt.Printf("週薪：%d\n\n", pay)
}

func payDay(hoursWorked, hourlyRate int) int {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrHourlyRate {
				fmt.Printf("時薪：%d\n錯誤：%v\n", hourlyRate, r)
			}
			if r == ErrHoursWorked {
				fmt.Printf("工時：%d\n錯誤：%v\n", hoursWorked, r)
			}
		}
		fmt.Printf("計算週薪的依據：工時：%d / 時薪：%d\n", hoursWorked, hourlyRate)
	}() //() 指立即執行匿名函式
	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime
	}
	return hoursWorked * hourlyRate

}
