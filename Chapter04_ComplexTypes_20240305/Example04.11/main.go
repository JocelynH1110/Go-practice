// 4-4-3  從切片和陣列建立新的切片

/*
取出一段"新切片"（最常見的做法）：
＊＊<新切片> = <陣列或切片>[<起始索引>:<結束索引（不含）>]

這會讓 Go 語言參考來源陣列或切片，將指定範圍的元素放進新切片中。
這範圍會從起始索引算起、一直到結束索引的前一個索引。
起始和結束索引都非必要，
若省略起始，會從第一個元素（索引 0 ）開始取值。
若省略結束，會取道最後一個索引（長度減 1）。
若兩個都省略（寫成：），就等於是取出原集合的所有元素。

使用上方的方式建立新切片時，Go 語言並不是真的將資料複製到新切片。
*/

// 練習、從切片再建立其他切片
package main

import "fmt"

func message() string {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	m := fmt.Sprintln("第一個元素：", s[0], s[0:1], s[:1])
	m += fmt.Sprintln("最末的元素：", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:])
	m += fmt.Sprintln("前五個元素：", s[:5])
	m += fmt.Sprintln("末四個元素：", s[5:])
	m += fmt.Sprintln("中間五個元素：", s[2:7])
	m += fmt.Sprintln("全部的元素：", s[:])
	return m
}
func main() {
	fmt.Println(message())
}

//通常真實世界的程式碼都只須處理一小部份的切片或陣列，而且也可以用 for range 迴圈來走訪每個元素。
//如果產生新切片時，起始和結束索引都不寫，就等於是將整個陣列轉換成切片。
//但這方式沒辦法拿來複製另一個切片，因為這會使兩個切片共享相同的底層資料。