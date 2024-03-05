// 4-3-5  寫入值到陣列

/*
只要陣列定義完畢，就可以用索引修改各別元素，這種賦值方式跟核心型別變數的做法一樣

＊＊<陣列>[<索引>]=<值>

*/

package main

import "fmt"

func message() string {
	arr := [4]string{"readt", "Get", "Go", "to"}
	arr[1] = "It's" //改變元素值
	arr[0] = "time"

	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}

func main() {
	fmt.Print(message())
}
