// 2-3  switch 敘述
// 2-3-1  switch 敘述基礎
/*
雖然在 if 敘述中，要加多少 else if 都可以，但過多會顯得雜亂、難以閱讀。
此時可以引用 GO 語言的另一種條件判斷敘述： switch 。

＊＊語法如下：
switch	<起始賦值敘述>; <運算式> {
case	<運算式>:
		<程式敘述>
case	<運算式>:
		<程式敘述>
		fallthrough
...
default:
		<程式敘述>
}

if 敘述只能使用布林運算式，switch 的運算式能做很多，回傳值不只能是布林值而已。
switch 的起始賦值和運算值都非必要，可以單寫一個，也可以兩個都不寫。若沒有運算式，效果就跟寫成 switch true 是一樣的。

case <運算式> 有兩個寫法：
1.寫成 if 敘述那樣的布林運算式。
2.直接寫一個值。
當 case 的值和 switch 值相符，其下的程式碼就會被執行。

GO 語言會由上往下檢查 case 的值或運算式。當找到時就會執行他的程式敘述並離開 switch 。
這和其他語言不同，C 語言的話要寫 break;才會跳出，不然就會全部都找並將符合的印出來。

fallthrough 用在 case 區塊中，這時不管下一個 case 條件是否符合，都會執行該區塊內容，跟在 C++ 中不使用 break; 一樣。

default 在所有 case 的運算式都不成立時，就會執行。它也可以省略，能放在 switch 的任何位置，但習慣上是放在最後面。
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Weekday(6)
	switch day {
	case time.Monday:
		fmt.Println("星期一，猴子穿新衣")
	case time.Tuesday:
		fmt.Println("星期二，猴子肚子餓")
	case time.Wednesday:
		fmt.Println("星期三，猴子去爬山")
	case time.Thursday:
		fmt.Println("星期四，猴子去考試")
	case time.Friday:
		fmt.Println("星期五，猴子去跳舞")
	case time.Saturday:
		fmt.Println("星期六，猴子去斗六")
	case time.Sunday:
		fmt.Println("星期日，猴子過生日")
	default:
		fmt.Println("日期不正確")
	}

}
