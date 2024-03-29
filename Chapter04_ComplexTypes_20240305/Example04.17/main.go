// 4-6   簡易自訂型別（custom types）
/*
＊＊建立自訂型別
  type <自訂型別名稱> <核心型別>

例、以字串為基礎建立一個叫 id 的零別
	type id string

自訂型別的行為和其核心型別一樣，包括擁有零值、能和同型別的資料對比等。
但自訂型別不能直接和它根據的核心型別相互做比較，除非先轉換其型別。

自訂型別的重點：可以加上自訂的行為（函式或方法），核心型別則無法。
*/
package main

import "fmt"

type id string

func getIDs() (id, id, id) {
	//用自訂型別建立變數
	var id1 id
	var id2 id = "1234-5678"
	var id3 id
	id3 = "1234-5678"
	return id1, id2, id3
}
func main() {
	id1, id2, id3 := getIDs()
	fmt.Println("id1==id2:", id1 == id2)
	fmt.Println("id2==id3:", id2 == id3)
	fmt.Println("id2 == \"1234-5678\":", string(id2) == "1234-5678")
}

//補充、型別別名
/*
Go 語言允許對型別取別名（alias），這樣並不會創造出新型別，而是能用不同的名稱來使用該型別：
＊＊ type <別名> = <型別>

例、定義 type num = int ，會替 int 型別創造一個別名 num ，於是你能將 num 當成 int 使用，而它仍然會被視為 int 型別。
*/
