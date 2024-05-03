// 6-4-3  建立 error 值
/*
在 error.go 裡有一個函式，可以用來建立自己的 error 值：
func New(text string) error {
	return &errorString{text}
}

Error() 函式會接收一個字串引數，並以此產生一個新的指標結構變數 *errorString ，再以 error 介面型別傳回。
雖說傳回值是 error 介面，但實際上傳回的是 *errorString 型別。
以下為證明程式碼：

package main

import (
	"errors"
	"fmt"
)

func main() {
	ErrBadData := errors.New("some bad data")
	fmt.Printf("ErrBadData type: %T\n", ErrBadData)
}

errors.New() 能讓你快速產生包含自訂訊息的 error 值，無須自己另外定義一個符合 error 介面的型別。
在 Go 語言中，error 值的名稱習慣上會以 Err 開頭（首字大寫），並採用駝峰式命名法。
*/

//練習、建立一個週薪計算程式：
/*
此函數會接收兩個引數，一個是該周工作時數，一個是時薪。
函式也要檢查這兩個參數是否有效，並且計算加班費。
1.時薪 10~75 美元
2.一週工時 0~80 小時
3.工時超過 40 小時，額外工時的時薪成以 2
4.若時薪或工時有誤，週薪傳回 0 ，並傳回對應的 error 值。若沒有錯誤，傳回計算後的週薪，錯誤則傳回 nil
*/
package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("無效的時薪")
	ErrHoursWorked = errors.New("無效的一週工時")
)

func payDay(hoursWorked, hourlyRate int) (int, error) {
	if hourlyRate < 10 || hourlyRate > 75 {
		return 0, ErrHourlyRate
	}

	if hoursWorked <= 0 || hoursWorked > 80 {
		return 0, ErrHoursWorked
	}
	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		hoursRegular := hoursWorked - hoursOver
		return hoursRegular*hourlyRate + hoursOver*hourlyRate*2, nil
	}
	return hoursWorked * hourlyRate, nil
}

func main() {
	//工時不正確
	pay, err := payDay(81, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)

	//時薪不正確
	pay, err = payDay(80, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)

	//加班費的計算
	pay, err = payDay(80, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)

	//週薪計算
	pay, err = payDay(40, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)
}
