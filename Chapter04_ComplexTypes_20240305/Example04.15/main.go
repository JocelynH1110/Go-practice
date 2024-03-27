// 4-5-2 從 map 從讀取元素
/*
只有在用索引鍵在 map 中取資料值時，才會知道該鍵存在與否。當索引鍵不存在，Go 語言就會傳回 map 資料值型別的零值。
藉由檢查零值來判斷索引鍵是否存在也是可行，但沒辦法永遠擔保零值就表示查無此鍵（零值也可能是有意義的資料）。

＊＊在這情況下可以在取值時，在 map 多接收一個參數，如下：
<值>, <存在狀態> := <map名稱>[<索引鍵>]

存在狀態是布林值。若 map 含有你傳入的索引鍵，該值就會為 true，反之 false 。
*/

// 練習、讀取 map 元素並檢查它存在與否
package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	return map[string]string{
		"305": "Susan",
		"408": "Tidi",
		"604": "Gigi",
	}
}

func getUser(id string) (string, bool) {
	users := getUsers()
	user, exists := users[id]
	return user, exists
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("未傳入使用者 ID")
		os.Exit(1)
	}
	userID := os.Args[1]
	name, exists := getUser(userID)
	if !exists {
		fmt.Printf("查無傳入的使用者 ID (%v).\n使用者列表：\n", userID)
		for key, value := range getUsers() {
			fmt.Println("使用者 ID :", key, "名字：", value)
		}
		os.Exit(1)
	}
	fmt.Println("查得名字：", name)
}

//os.Exit：用於中止程序的執行，並返回一個特定的退出碼。os.Exit(0)：表成功，非零代表錯誤。
//panic：終止當前正在運行的程式（包括所有協程）並輸出導致異常的堆棧信息。

/*
你所看到的輸出結果不一定跟課本範例一樣順序，因為 Go 語言會在你使用 for range 時故意打亂元素。
從上練習可以學到，如何檢查 map 中某個索引鍵是否存在。
別種程式語言都會要求在取值前先確認索引鍵是否存在，Go 則先取值再檢查，但這方式可以大幅減少執行期間錯誤。

若你的程式無法用資料零值來判別鍵是否存在（map 找不到索引鍵時就會傳回零值），就可利用 map 的第二個傳回值來檢查索引鍵是否存在。
*/
