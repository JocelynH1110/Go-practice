// 4-5-3  從 map 刪除元素
/*
從 map 移除元素，作法和陣列和切片都不一樣：
陣列：元素無法移除，因為長度已經固定，頂多只能把該元素設為零值。
切片：可以用零值清空該元素，或是用 append() 組合新的切片範圍，並在過程中去掉某些元素。
map：可以將元素變為零值，但元素仍然存在，所以程式在檢查時會得到「鍵仍然存在」的錯誤結論。也無法像切片一樣靠擷取範圍來「丟掉」元素。

＊＊移除 map 元素，必須引用 Go 語言的內建函式 delete()：
delete(<map 元素>, <索引鍵>)

delete() 沒有任何傳回值：如果要刪除的索引7不存在，它也不會發生問題或提出異議。
*/

// 練習、從 map 刪除一個元素
package main

import (
	"fmt"
	"os"
)

var users = map[string]string{
	"101": "Ruby",
	"201": "Luby",
	"301": "Puby",
	"401": "Tuby",
}

func deleteUser(id string) {
	delete(users, id)
}
func main() {
	for len(os.Args) < 2 {
		fmt.Println("未傳入使用者 ID")
		os.Exit(1)
	}
	userID := os.Args[1]
	deleteUser(userID)
	fmt.Println("使用者列表：", users)
}

//此內建函式 delete() 唯一限制是只能搭配 map ;不能對陣獵獲切片使用。
