// 7-3 鴨子定型和多型
// 7-3-1 鴨子定型（duck typing）也稱鴨子型別
/*
鴨子定型是程式設計中的一種歸納理論：「只要一個東西長的像鴨子、游泳像鴨子、叫聲像鴨子，那它就是鴨子。」
以 Go 語言來說，任何型別只要符合某個介面的行為規範，那他們就通通能當成該介面型別來使用。

意即，Go 語言的鴨子定型是根據型別方法來判斷型別符合介面，而不是明確地指定哪些型別能夠符合。
下面來看個例子：
*/
package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct {
}

func (c cat) Speak() string {
	return "PurrrrMeow"
}

func chatter(s Speaker) { //接收 Spraker 介面型別的引數
	fmt.Println(s.Speak())
}

func main() {
	c := cat{}
	chatter(c)
}

//我們有個函式 chatter()，它接收的參數型別是 Speaker 介面。cat 結構隱性實作了 Speaker 介面，因此透過鴨子定型被視為 Speaker 型別，可以傳入 chatter() 的參數。
