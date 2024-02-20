//1-8	列舉（enums）
/*
列舉是一種定義一系列常數的方式，這些常數的值是整數，而且彼此相關。
GO 語言沒有內建列舉專用的型別，但提供了一種稱為 iota 的工具，讓你可以用常數定義出自己的列舉資料。
*/

// 例子、我們要在以下程式碼將一週中的每一天定義為常數：
package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuseday
	Wednesday
	Thursday
	Friday
	Saturday
	ooo
)

func main() {
	fmt.Println(Monday)
	fmt.Println(Friday)
	fmt.Println(ooo)
}

//iota 會指派數值
