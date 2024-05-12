// 7-3-2 多型（polymophism）
/*
多型：指一樣東西可以用多種形式呈現。例如、一個形狀可以是正方形、圓形、矩形，或是任意其他形狀。

在其他物件導向程式語言中，子類別化（subclassing）意指讓一個類別繼承（inherit）另一個類別的欄位和行為。若設計出多個子類別，每個子類別都經過修改而各有差異，這就是物件導向中的多型。 Go 不是物件導向語言，它沒有類別，但仍可以透過內嵌結構和介面來實現類似子類別化的概念。

在 Go 語言使用多型的好處之一：若首上有寫好且經測試的程式碼，就可重複利用它。只要讓該函式接收介面型別參數，那麼任何符合介面規範的型別都可以傳入，而不是只限於核心型別int、float、bool。甚至不需在函式中撰寫額外的程式碼來應付每一種型別，畢竟只有正確實作介面的型別才有辦法傳入你的函式。

任何實質型別都能實作一種以上的介面，反之，同一個介面也能被多個型別實作。例如、Speaker 可以同時由 dog、cat、person 型別實作。

如果 cat、dog、person 都實作了 Speaker 介面，這代表他們一定都有 Speaker() 方法，且會傳回一個字串。這表示可以撰寫一個共同函式，接收 Speaker 介面型別的參數，然後對任何傳入的值呼叫相同的行為：（以下為例子）
*/
package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct {
}

type dog string // dog 自訂型別

type person struct { // person 結構
	name string
	age  int
}

func main() {
	c := cat{}
	d := dog("")
	p := person{name: "Heather", age: 20}
	thingSpeak(c)
	thingSpeak(d)
	thingSpeak(p)

	//改寫成接收數量不定的參數後可以縮寫成：
	thingSpeak(c, d, p)
}

func (c cat) Speak() string {
	return "Purr Meow"
}
func (d dog) Speak() string {
	return "Woof woof"
}
func (p person) Speak() string {
	//return "Hi,my name is " + p.name + "." + p.age
	return fmt.Sprintf("Hi,my name is %v %v", p.name, p.age)
}

/*
func thingSpeak(s Speaker) {
	fmt.Println(s.Speak())
}
*/

// 把 thingSpeak() 換成接收數量不定的參數（lesson 5）
func thingSpeak(speakers ...Speaker) {
	for _, s := range speakers {
		fmt.Println(s.Speak())
	}
}

// fmt 套件的 Print()/Println() 也是這樣接收多重參數的
