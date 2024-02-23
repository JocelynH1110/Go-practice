// 2-2-4	if 敘述的起始賦值
/*
我們常會呼叫某個函式，但只會拿函是的回傳值來檢查它是否正確執行。之後就在也不需要這個值。
在這種情況下，函式傳回的變數雖然後面根本用步道，但只要你接收它，這些資料就仍會存在於其作用範圍內，等於多佔了一份記憶體。
為了避免這種浪費，可以把這些變數的作用範圍限制在 if 敘述範圍，這樣只要一離開 if 敘述，該變數就會消滅。
＊＊方法就是在 if 敘述中加上，起始賦值敘述（init statement）＊＊
If <起始賦值敘述>; <布林值運算>{
	<程式區塊>
}

以分號做區隔，布林運算式可以直接使用起始賦值敘述內宣告的變數來做判斷。
起始賦值敘述只能用以下方式做簡單敘述：
1.短變數宣告。例、i:=0	(不能使用 var)
2.算術或邏輯運算式。例、i:=(j*10)==40
3.遞增或遞減運算式。例、i++
4.在並行性運算中傳值給通道的敘述（lesson 16）
*/
package main

import (
	"errors"
	"fmt"
)

func validate(input int) error {
	if input < 0 {
		return errors.New("輸入值不得為負")
	} else if input > 100 {
		return errors.New("輸入值不得超過 100")
	} else if input%7 == 0 {
		return errors.New("輸入值不得為 7 的倍數")
	} else {
		return nil
	}
}
func main() {
	input := 21
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "是偶數")
	} else {
		fmt.Println(input, "是奇數")
	}
}

//一旦 main() 的敘述完成任務，則 err 變數離開作用範圍，被 GO 語言的記憶體管理系統回收。
