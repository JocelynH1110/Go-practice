// 7-4-4 型別斷言、型別 switch

// 型別 switch
/*
當空白型別背後的實質型別有很多種，無法事先知道是什麼，未避免寫下眾多型別斷言，型別 switch 就派上用場。

＊＊＊型別 switch 語法：
switch v:=i.(type){
case S:
	// v 是型別 S 時要執行的程式碼
}

switch 和斷言十分類似，只差在把型別換成關鍵字 type 而已。
型別 switch 會比對每個 case 後面的型別，找尋吻合的對象。
*/
package main

import "fmt"

type cat struct {
	name string
}

func main() {
	// 建立一個空介面切片，放入不同型別的值
	i := []interface{}{42, "The apple", true, cat{name: "Rara"}}
	typeExample(i)
}

func typeExample(i []interface{}) {
	for _, x := range i {
		switch v := x.(type) {
		case int:
			fmt.Printf("%v is int\n", v)
		case string:
			fmt.Printf("%v is string\n", v)
		case bool:
			fmt.Printf("%v is bool\n", v)
		default:
			fmt.Printf("%T is unknown type\n", v)
		}
	}
}
