// 4-3-4  讀取陣列元素

/*
＊＊讀取陣列中的單一元素：
<值> = <陣列>[<索引>]

陣列中元素的位置是保證固定的。表示只要元素放在某個索引鍵的位置，它就永遠能用那個索引鍵存取。

這種用位置對應資料的特性，在試算表很常見的 CSV （comma-separated value，讀取逗號分隔點）或其他用分隔符號區分資料的檔案格式都很有用
*/

package main

import "fmt"

func message() string {
	arr := [...]string{
		"ready",
		"get",
		"go",
		"to",
	}
	//用 fmt.Sprintln() 傳回格式化字串
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}
func main() {
	fmt.Print(message())
}
