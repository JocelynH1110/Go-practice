// 6-5-2 panic() 函式
// panic 時 defer 的執行效果

// 例子、發生 panic 後的程式碼就不會執行了，來看看若 panic 後面的函式加上 defer 敘述（lesson 5）會如何操作：
package main

import (
	"errors"
	"fmt"
)

func main() {
	defer fmt.Println("在 main() 使用 defer")
	test()
	fmt.Println("這一行不會印出")
}

func test() {
	defer fmt.Println("在 test() 使用 defer")
	msg := "bye"
	message(msg)
}

func message(msg string) {
	defer fmt.Println("在 message() 使用 defer")
	if msg == "bye" {
		panic(errors.New("又出事惹"))
	}
}

// 結果會是印出 message()->test()->main() 的 defer 最後 panic 裡的錯誤訊息。

/*
Note、
os.Exit() 會立即中斷程式，並傳回一個狀態碼（習慣上 0 代表正常，正整數代表錯誤，其意義由使用者決定）
但使用 os.Exit() 時一律不會執行被延遲的函式。
因此在特定情況下，使用 panic() 會比 os.Exit() 來得合適。
*/

// 練習、利用 panic() 函式讓程式在發生錯誤時當掉
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
	pay := payDay(81, 50)
	fmt.Println(pay)
}

func payDay(hoursWorked, hourlyRate int) int { //不回傳 error
	//不管有無引發 panic ，在 payDay() 結束時印出工時與薪資
	report := func() { //匿名函式
		fmt.Printf("工時：%d\n時薪：%d\n", hoursWorked, hourlyRate)
	}
	defer report()

	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}
	return hoursWorked * hourlyRate
}
