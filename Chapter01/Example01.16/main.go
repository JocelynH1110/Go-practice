// 1-6-3	 從指標取得值
package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Timer{}

	if count1 != nil {
		fmt.Printf("count1:%#v\n", *count1) //用 * 取得指標的值
	}

	if count2 != nil {
		fmt.Printf("count2:%#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3:%#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("time:%#v\n", *t)

		//		fmt.Printf("time:%#v\n", t.String())
	}

}

//以上例子用解除參照來從指標取出實際值，同時加上 nil 檢查，以免遇到解除參照錯誤。
