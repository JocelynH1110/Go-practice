// 1-3-2 一次更改多個變數值
package main

import "fmt"

func main() {
	//宣告多重變數並賦初始值
	query, limit, offset := "bat", 10, 0

	//用單行敘述一次更改所有變數的值
	query, limit, offset = "ball", offset, 20

	fmt.Println(query, limit, offset)
}

/*補充：
正常情況下，不能對已宣告過得變數使用短變數宣告來賦值。
例外、若短變數宣告左側有多重變數，其中有一變數是之前沒有的，這樣就可以成立。
*/
