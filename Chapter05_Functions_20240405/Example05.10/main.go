// 5-6  defer
// 5-6-1 用 defer 延後函式執行
/*
defer 敘述能延後函式的執行時機，使該函式等到父函式結束（跑完所有程式碼或執行 return ）的前一刻才會被執行。

被延後的函是有什麼好處呢？
通常是用來「善後」的，包括像是釋出資源、關閉已開啟的檔案、關閉仍在連結的資料庫連接、移除程式先前建立的設定/暫存檔案等。
此外、defer 函式也可用來從程式的錯誤狀況復原（chapter 6）
*/

//例子、用具名函式寫。
//defer 延後了 done() 函式的執行時間
/*
package main

import "fmt"

func main() {
	defer done()
	fmt.Println("main() 開始")
	fmt.Println("main() 結束")
}

func done() {
	fmt.Println("換我結束了")
}
*/

// 例子、用匿名函式寫。
package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("換我結束了")
	}()
	fmt.Println("main() 開始")
	fmt.Println("main() 結束")
}

//從某種程度上來說，用具名函式的寫法有較好的可讀性。
