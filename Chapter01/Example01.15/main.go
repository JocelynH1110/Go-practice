// 1-6	值 vs. 指標（pointers）
// 1-6-1  了解指標
/*
當我們把 int、bool、string 這類值傳給函式處理時，GO 語言會在函式中複製這些值，建立出新的變數。
這複製的動作意味著你呼叫函式時，若函式對參數做出更動，原始值也不會受影響，能減少程式碼的錯誤。
這種傳值的方式 GO 語言採用了一種簡單的記憶體管理系統叫 "堆疊" ，每個參數都會在堆疊中獲得自己的記憶體。
缺點是越多值在函式間傳遞，這複製動作會消耗越多的記憶體。到頭來可能會消耗比實際需求還多的記憶體

另一種函式傳值的替代方式，用的記憶體較少。其不會複製值，而是建立指標在傳遞給函式。
指標跟值本身是兩回事，而指標唯一用途就只是拿來取得值而已。
在使用指標傳值給函式時，GO 語言就不會複製指標指向的值。
用指標時 GO 語言會把值放在所謂的堆積（heap）記憶體空間：允許一個值存在，直到程式中沒有指標參照到它為止。
GO 語言有垃圾回收機制（garbage collection）程序來回收這些記憶體，這機制會在背景定期運作。
缺點是當堆積內累積了大量資料時，回收機制就必須做大量的檢查，這會耗損 CPU 週期。

除了改善效能的目的，也能用指標改變程式設計的方式。
例如當想判斷某值是否存在，判斷一般的變數就會有問題，因為它一定至少會帶有零值，其在程式中可能仍然合法。
相對的，指標有未設定（is not set）的狀態，當它沒有儲存目標值的位址時會傳回 nil ，nil 在 GO 語言中代表無值。

指標本身可以是 nil 的這種特性，意味著就算指標沒有指向任何值，還是可以取得指標本身的值，進而導致執行期間錯誤（runtime error）。
為避免，可以先拿指標和 nil 做比較（<指標>!= nil），然後才對指標賦值。

*/

// 1-6-2	取得指標
/*
取得指標的方式：
1.把型別設為指標（這方式其初始值會是nil）
var	<變數>	*<型別>

2.內建函式 new() 可以達到賦值效果。
該函式用意在於為某種型別取得記憶體、填入該型別的零值，後回傳該記憶體指標。
<變數> := new(<型別>)
var	<變數> = new(<型別>)

3.取得某個既有變數的指標，用＆算符
<變數 1> := &<變數 2>
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Timer{}

	fmt.Printf("count1:%#v\n", count1)
	fmt.Printf("count2:%#v\n", count2)
	fmt.Printf("count3:%#v\n", count3)
	fmt.Printf("time:%#v\n", t)
}