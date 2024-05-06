// 7-1 前言
/*當函式解析和處理資料的行為完全一樣，只有讀入的資料型別有所不同，那能不能只寫一個函式就好呢？又該如何繞過不同資料型別所帶來的限制，讓同一個函式能接收不同型別的值呢？

Answer：使用介面。可以寫一個 loadEmployee(r io.Reader) 來取代以上三個函式。Go 語言內建的 io.Reader 型別就是所謂的介面，可以接受包括 string、os.File、http.Request 在內的不同型別。

在本章中，將來研究介面為何能接收多重型別，並探討介面如何引進「鴨子定型」（duck typing）、「多型」（polymorphism）的機制。
最後會回顧第四章的空介面型別，及他要怎麼搭配型別斷言、型別 switch 來檢查空介面底下的型別。
*/

// 7-2 介面（interface）
// 7-2-1 認識介面 +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/* Go 語言的介面型別會包括一系列函式或方法特徵（method signatures），而其他型別只要定義了完全相同的方法，就等於是實作（implement）了此介面、並可被當成該介面型別來看待：

＊＊type 型別名稱 interface{
	<方法 1 特徵>
	<方法 2 特徵>
	...
}

// 7-2-2 定義介面型別 ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/*
＊＊定義介面的步驟：
type Speaker interface{
	Speak(message string) string	//()參數、傳回值
	Greet() string
}

1.關鍵字 type、介面名稱（動詞）、關鍵字 interface。
2.慣例上介面名稱會拿其中一個方法的名稱加上 er 結尾，特別是介面中就只有一個方法的時後。
3.在大括號內定義方法特徵。

任何型別要符合這個介面，就得實作出上面的方法（所有方法的特徵也都得符合才行）。
總結，介面是一種型別，但其內容就是方法特徵的集合。
*/

// 7-2-3 實作一個介面 ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/*
其他語言都得以明確的方式實作介面，意思是必須明白陳述一個物件是在沿用哪個介面的規範。
Java 來說：
class Dog implements Pet

以上指出，Dog 類別（class）實作了 Pet 介面。它就必須實作介面要求的方法，否則會產生錯誤。
然而在 Go 語言中，介面實作是隱性的。只要一個型別綁定的方法特徵完全符合一個介面規範，該型別就等於是自動實作了該介面。
*/
package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type cat struct {
}

func main() {
	c := cat{}
	fmt.Println(c.Speak())
	c.Greeting()
}

func (c cat) Speak() string { // cat 的方法（這使 cat 符合 Speaker 介面)
	return "Purrr meow"
}

func (c cat) Greeting() { // 這方法並未定義在 Speaker 中，但既然 cat 型別已經滿足了 Speaker 型別，這便不影響 cat 對 Speaker 介面的實作。
	fmt.Println("Meow,mmmmeeeooowwww!!")
}

//程式裡並沒有明確的敘述指出 cat 是在實作 Speaker 介面；當你替 cat 定義 Speaker() 方法時，隱含實作就發生了。
