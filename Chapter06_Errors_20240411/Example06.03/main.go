// 6-4 error 界面
// 6-4-1  Go 語言的 error 值
/*
在 Go 語言中，一個 error 是一個值。

引言、資料值是能用程式設定的，而既然 error 也是，所以 error 同樣可以用程式設定內容。error 不是程式例外，它並無任何特別之處，因為未處理的例外會讓程式當掉，error 卻不見得。

既然 error 是個值，它就可以被當成引數傳給函式、被函式傳回，並能像 Go 語言的任何值一樣被讀取和做比較。

事實上，Go 語言的 error 值都必須實作來符合 error 介面的定義（lesson 7）。
＊＊下面是 Go 語言中宣告的 error 介面型別：
type error interface{
	Error() string
}

NOTE、這個型別被宣告在最高層級，因此任何 Go 程式都能直接取用。

一個型別只要符合介面的規範（擁有一樣的方法函式），就會被視為符合該介面型別。
任何型別只要符合 error 介面的要求，它就能當 error 型別：
1.型別得擁有一個函式（或稱方法）叫 Error()
2.Error() 得傳回一個 string 型別的值

在 GO 語言標準函式庫中，各套件可能會定義自己的 error 型別，包含不同欄位和方法。
但只要這些型別具備 Error() string 方法，就可以用 error 介面型別的形式建立 error 值，且能被許多跟錯誤處理有關的功能共用。
*/

// 例子、以程式內的錯誤處理部份來說明 Go 語言如何處理錯誤：
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "10"
	s, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T,%v\n", s, s)

	v = "s2"
	s2, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T,%v\n", s2, s)
}

//任何函數若要傳回錯誤，error 就應該當成最後一個傳回值。
//以上例子重複使用 err 值是 Go 語言常見的撰寫習慣，畢竟在檢查過他的值後就用不到了。

/*
若 error 值不是 nil 代表有錯誤發生，必須決定如何應之。依據不同場合，想做的事可能如下：
1.將 error 傳給函式呼叫者
2.用 log 紀錄錯誤然後繼續執行
3.停止程式執行
4.忽略 error (極度不建議)
5.引發 panic (只有罕見的狀況才會這樣做)
*/
