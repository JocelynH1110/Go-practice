//1-9	變數作用範圍（Scope）
/*
在 GO 在語言中，所有的變數都有其運作範圍。最頂層的範圍是套件（package）範圍。
子範圍定義，最簡單的辨識方式就是{}內。
當某段程式碼存取某個變數時，GO 語言會檢查該程式碼的運作範圍，若在該範圍找不到該名稱，就會"往上一層範圍"找。
若是同名變數，但型別不同，也會拋出錯誤訊息。
*/

// 練習、從子範圍存取上層變數
package main

import "fmt"

var level = "pkg" //套件範圍變數

func main() {
	fmt.Println("Main start：", level) //main() 層級
	if true {
		fmt.Println("Block start：", level) //main() 底下的 if 層級
		funcA()
	}
}

func funcA() {
	fmt.Println("funcA start：", level) //funcA() 函式層級
}
