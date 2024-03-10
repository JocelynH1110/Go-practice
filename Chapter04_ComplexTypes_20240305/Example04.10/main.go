// 4-4-2  為切片附加多重元素
/*
append() 可以一次附加多個值到切片裡，因為該函式的第二個參數可以接收數量不定的值：
＊＊<新切片> = append(<切片>, <新元素 1>,<新元素 2>,<新元素 3>...)

在 append() 傳入一個切片，並在後面加上解包算符（unpack operator），也就是三個點，來解開它，使切片的元素會被拆成單獨的值傳入 append()，有多少元素就傳多少。
真實世界很常用這種方式，把一個以上的參數傳給 append() 處理。保持程式精簡，不一定得用迴圈來新增值到切片。
*/

// 練習、一次為切片加入多個新元素
package main

import (
	"fmt"
	"os"
)

func getPassedArgs() []string {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func getLocals(extraLocals []string) []string {
	var locals []string

	locals = append(locals, "en_US", "fr_FR") //加入預設元素

	locals = append(locals, extraLocals...) //加入使用者提供、數量不定的參數。沒有這行的話，輸入的參數不會解包加入。
	return locals
}

func main() {
	locals := getLocals(getPassedArgs())
	fmt.Println("要使用的語系：", locals)
}
