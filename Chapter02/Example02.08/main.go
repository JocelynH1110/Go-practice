// 2-4-3  for range 迴圈
/*
像陣列和切片這樣的集合，一定會有連續索引值存在，並從 0 開始計算，前面介紹的 for i 迴圈，就是處裡這類資料最常用的工具。

另一種形式的資料集合 map ，其鍵與值不會照順序排列，對它使用 for i 迴圈就沒有那麼便利。
這表示我們必須使用 range 來取代原本迴圈裡原本的條件敘述。
range 每次會從集合取出一個鍵與值，下一輪迴圈執行時就換下一組。
使用 range 敘述時，就不需要定義 for 迴圈的結束條件，range 會自己處理好（在取完集合內所有值後結束迴圈）
*/

package main

import (
	"fmt"
)

/*1.22 版本後可以用數字直接寫，不然都要寫i>0 之類的
func main() {
	for i := range 20 {
		fmt.Println(i)
	}
}
*/

// 練習、利用迴圈走訪 map 元素
func main() {
	config := map[string]string{ //建立 map ，元素由一對對鍵與值構成。
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}

	for key, value := range config {
		fmt.Println(key, "=", value)
	}
}

//補充、
//如在迴圈中用不到 key、value 變數，可在接收時寫成底線字元_，來告知編譯器說你不需要它
/*	1.for _, value := range config {
		fmt.Println(key, "=", value)

	2.for key:= range config {
		fmt.Println(key, "=", value)
	  for key, _:= range config {
		fmt.Println(key, "=", value)
*/

/* range 敘述也可用於陣列和切片，這情況下 key 會是元素索引，value 會是元素索引，則是元素值
	names:=[]string{"Cece","Fifi","Dada"}

	for i ,value:= range names{
		fmt.Println("Index",i, "=" ,value)
	 }

但若你要在迴圈中修改原始集合內的元素，就得使用 name[i]，因為 value 是個在迴圈內建立的獨立變數，和原集合沒有關係。
*/
