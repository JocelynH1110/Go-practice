// 2-2-3	 else if 敘述
/*
可以加入任何數量的 else if，GO 語言會由上往下依序檢視含有布林運算式的敘述，直到找到結果為true，然後只執行該程式區塊。
如果都沒有，就會執行最後的 else。
若是沒有最後的 else，前面程式布林值運算也沒有一個為 true，則 GO 就不會執行任何程式區塊。
語法如下：
if <布林值運算 1>{
	<程式區塊 1>
}else if <布林值運算 2>{
	<程式區塊 2>
}else if <布林值運算 3>{
	<程式區塊 3>
}else {
	<程式區塊 N>
}
*/
package main

import "fmt"

func main() {
	input := -10

	if input < 0 {
		fmt.Println("輸入值不得為負！")
	} else if input%2 == 0 {
		fmt.Println(input, "是偶數")
	} else {
		fmt.Println(input, "是奇數")
	}
}
