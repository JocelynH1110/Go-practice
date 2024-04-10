// 5-6-2  多重 defer 的執行順序
/*
可以在同一個函式內使用多個 defer 敘述來延後多個函式，但被延後的順序可能會跟想像的有所出入。
當我們對多個函式使用 defer 時，其執行順序會遵循「先進後出」（First In Last Out,FILO）。
*/

package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("我是第一個宣告的！")
	}()
	defer func() {
		fmt.Println("我是第二個宣告的！")
	}()
	defer func() {
		fmt.Println("我是第三個宣告的！")
	}()
	f1 := func() {
		fmt.Println("f1 開始")
	}
	f2 := func() {
		fmt.Println("f2 開始")
	}

	f1()
	f2()
	fmt.Println("main() 結束")
}
