// 5-2-4 Naked Returns
/*
如果沒有在 return 敘述後面指定要傳回的變數，Go 語言會將傳回值清單裡的變數傳回，這便是所謂的 naked return，也可以稱為具名 return （named return）：
func greeting()(name string,age int){
	name = "Fala"
	age = 30
	return  //傳回 name ,age 變數
}

naked return 的缺點之一：
若用在較長的函式中，可能會讓讀程式碼的人稿混，弄不清楚到底傳回了什麼變數。故在函式稍微複雜的情況下，應避免使用 naked return。

此外、naked return 還可能衍生出變數遮蔽（shadowing）問題：
func message()(message string,err error){
	message = "hi"
	if message =="hi"{
		err:=fmt.Errorf("say bye\n")
		return
	}
	return
}

以上程式碼會導致錯誤「err is shadowed during return」。
因為變數 err 先在函式的傳回值清單宣告過，接著又在 if 敘述的大括號範圍內被宣告和初始化，往上遮蔽（shadow）了函式層級的同名變數。
這時在 if 內部使用 naked return ，到底該傳回哪個 err 變數呢？Go 編譯器就會提出錯誤。
*/

// 練習、對應特定標頭的索引值：naked return 版
package main

import (
	"fmt"
	"strings"
)

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	result := csvHdrCol(hdr) //接收傳回值
	fmt.Println("Result:")
	fmt.Println(result)
	fmt.Println()

	hdr2 := []string{"Employee", "Empid", "Hours Worked", "Address", "manager", "Hourly Rate"}
	result2 := csvHdrCol(hdr2) //接收傳回值
	fmt.Println("Result2:")
	fmt.Println(result2)
	fmt.Println()
}

func csvHdrCol(hdr []string) (csvIdxToCol map[int]string) { //定義傳回值的名稱和型別
	csvIdxToCol = make(map[int]string) //初始化傳回變數
	for i, v := range hdr {
		switch v := strings.ToLower(strings.TrimSpace(v)); v { //用 TrimSpace() 把標頭去掉空白、ToLower()轉成小寫
		case "employee":
			csvIdxToCol[i] = v
		case "hours worked":
			csvIdxToCol[i] = v
		case "hourly rate":
			csvIdxToCol[i] = v
		}
	}
	return
}
