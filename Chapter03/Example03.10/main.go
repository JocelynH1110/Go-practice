// 3-5  nil 值

// nil 並不是一個型別，而是 Go 語言的一個特殊資料值，代表的是一個無型別也無值的狀態。
//在處理指標、map、介面（interfaces）及 error 值時，都必須確定他們不是 nil 。
//如果嘗試用一個 nil 值做運算，程式就會掛掉。

// 例、檢查資料值是否為 nil
package main

import "fmt"

func main() {
	var message *string //沒有初始值的指標變數會是 nil

	if message == nil {
		fmt.Println("錯誤，非預期的 nil 值")
		return
	}
	fmt.Println(&message)
}
