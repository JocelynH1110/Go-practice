// 7-2-4 隱性介面實作的優點
/*
當在其他程式語言中實作介面，必須明確表達意圖（比如在 Java 使用關鍵字 implement）。若修改了一個介面定義的方法，就得一一找出實作該介面的類別和修改、或將不符資格者的 implement 間鍵字移除。
在 Go 語言裡，只要型別的行為滿足了某個介面，就自動實作該介面，若後來修改了介面的方法集合，那麼不符合資格的實作也會自動失效（但這不會影響該型別的其他方面）。

另一個優點是，可以讓型別實作其他套件內定義的介面，這樣一來就能將介面與其實作型別的定義分開（去耦合，decouple）。（lesson 8）
*/

/*下面為一個在主程式（即 main 套件）中運用其他套件定義的介面的例子。fmt 套件的 Stringer 就是一個例子，不僅許多套件會用到，fmt 套件本身也會拿他來對主控制台印出資料：

type Stringer interface{
	String() string
}

以 fmt.Println() 為例，若傳入的型別符合 Stringer 介面型別，那麼 Println() 就會呼叫其 String() 方法來取得字串。
*/

// 例子、修改前面 cat 的結構範例，給它加上兩個欄位，刪除 Greeting() 方法，換上一個新方法 String():
package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct { //加入欄位
	name string
	age int
}

func (c cat) Speak() string {
	return "Purrr Meow"
}

func (c cat) String() string { // String() 方法
	return fmt.Sprintf("%v (%v years old)", c.name, c.age)
}

func main() {
	c := cat{name: "Oreo", age: 10}
	fmt.Println(c.Speak())
	fmt.Println(c) // 用fmt 套件直接印出 cat
}

//現在我們在 cat 結構型別同時實作了兩個介面，一個是在 main 套件內自訂的 Speaker，一個是來自 fmt 套件的 Stringer。
//目前我們還沒有程式會用到 Speaker 介面，但你會發現用 fmt.Println() 印出 cat 結構變數 c 時，Println() 自動呼叫了他的 String() 方法。這表示 c 符合並實作了 String 介面，使得它能夠被 Println() 接受和表現出特定的行為。

/*
補充、實作介面時使用「值接收器」和「指標接收器」的差別
若把上面程式碼改成，用指標接收器的形式來指向 cat 結構變數：
*/
func (c *cat) Speak() string {
	return "Purrr Meow"
}
func (c *cat) String() string {
	return fmt.Sprintf("%v (%v years old)", c.name, c.age)
}

/*以上例子可以看出 Speak() 依然正常運作，因為他是 cat 結構的方法，可 fmt.Println() 就沒有呼叫 cat.String() ，只單純印出結構內容。
這是因為加上指標接收器後，變成了指標型別 *cat 而不是型別 cat 實作了 Stringer 介面，使得在此建立的 cat 變數就不再被認為符合 Stringer 介面了。

解決方法是把 cat  變數宣告成指標：
	c := &cat{name: "Oreo", age: 10}

btw，方法若是使用接收器，該型別的變數不管是值或指標，都可以正確實作介面。
*/

// 練習、隱性實作一個介面 — 首先訂出一個 person 結構型別，含有 name、age、isMarried 等欄位。它擁有 Speak() 方法，隱含實作了我們自訂的 Speaker 介面，此外也有 String() 方法，好隱含實作 fmt 套件的 Stringer 介面。
package main

import "fmt"

type person struct {
	name      string
	age       int
	isMarried bool
}

type Speaker interface {
	Speak() string
}

func main() {
	p := person{name: "Tiaka", age: 27, isMarried: false}
	fmt.Println(p.Speak())
	fmt.Println(p)
}

func (p person) Speak() string {
	return "各位好，我的名字是" + p.name
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v 歲)\n已婚：%v", p.name, p.age, p.isMarried)
}
