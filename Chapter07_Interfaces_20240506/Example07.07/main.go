// 7-4-2 以介面為傳回值的函式
/*
func someFunc() Speaker{}{	//傳回值是 Speaker{} 介面
	// 程式碼
}

任何型別只要實作 Error() string 方法就能符合 Go 語言的 error 介面，而事實上每個套件都會定義他們自己的 error 型別。
*/

//例子、不同套件傳回的 error 值實際上是什麼型別：

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"` // 故意把欄位型別改錯
}

func main() {
	p, err := loadPerson("data.json") //讀取同目錄下的文字檔
	if err != nil {
		//若有錯誤，印出值和型別
		fmt.Printf("%v", p)
		fmt.Printf("%T", err)
	}
	fmt.Printf("%#v\n", p)
}

func loadPerson(fname string) (Person, error) {
	var p Person
	f, err := os.Open(fname)
	if err != nil {
		return p, err //傳回檔案開啟錯誤
	}
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		return p, err //傳回 JSON 解析錯誤
	}
	return p, err

}

// 究竟該不該使用介面為傳回值？
// 諺語：接收介面、傳回結構。函式接收介面可以增加使用者的實作彈性，但傳回值用介面會造成使用者實作上的混搖。使用者可能得做額外的型別斷言，更有可能花時間查詢程式文件，才能了解不同型別的欄位和其行為差異。

// 以下為簡單方針，協助判斷使用介面傳回值的場合：
/*
1.如果沒有絕對必要，就不要在函式傳回介面型別。
2.介面定義越精簡越好，好讓使用者更容易實作它。
3.盡量在實質型別（需求）存在後，才根據他們撰寫介面，而非反過來。
4.通常介面會定義在用到該型別的套件內（例如將它用於函式參數型別的套件）。對外用不到的介面就不應該匯出。
*/
