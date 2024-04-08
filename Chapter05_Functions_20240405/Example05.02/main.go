// 5-2-2 函式參數
/*
參數決定了你能把哪些引數（arguments）或值傳給函式。
函式可以沒有參數、也可以有多個參數。
但不應該弄出一大串參數，那只會程式碼更難讀。

選擇參數：只有函式的單一職責解決問題時需要的參數，才應該列入定義。

參數和引數的差別：
當函式定義為 greeting(name string,age int)
參數 → name 、age。
引數 → 當呼叫函式 greeting("Sala",30)，括號裡的字串和數字就是引數。

若用變數當引數，變數名稱與參數名稱也不須一致。
不管是什麼名稱的變數傳入函式，只要型別正確，其值就會被賦予給參數。

函式參數就是它的區域變數，亦即作用範圍只限函式內部，在函式以外就無法存取。
在呼叫函式時，引數傳入的型別與順序必須呼應參數的定義。
*/

// 練習、對應特定標頭的索引值：建立另一個函式，接收的參數是一份csv資料的標頭所構成的切片，要找尋特定標頭和他們所在的索引，並以 map 形式印出。
package main

import (
	"fmt"
	"strings"
)

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}

	csvHdrCol(hdr)

	hdr2 := []string{"Employee", "Empid", "Hours Worked", "Address", "manager", "Hourly Rate"}

	csvHdrCol(hdr2)
}

func csvHdrCol(header []string) {
	csvHeadersToColumnIndex := make(map[int]string)
	for i, v := range header {
		switch v := strings.ToLower(strings.TrimSpace(v)); v { //用 TrimSpace() 把標頭去掉空白、ToLower()轉成小寫
		case "employee":
			csvHeadersToColumnIndex[i] = v
		case "hours worked":
			csvHeadersToColumnIndex[i] = v
		case "hourly rate":
			csvHeadersToColumnIndex[i] = v
		}
	}
	fmt.Println(csvHeadersToColumnIndex)
}

//函式無法存取呼叫它的父函式的變數。若要存取那些變數，唯一方式就是透過參數傳遞。
