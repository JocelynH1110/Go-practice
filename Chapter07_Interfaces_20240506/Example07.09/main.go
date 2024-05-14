// 7-4-4 型別斷言、型別 switch

// 型別斷言（type assertion）
/*
讓你可以檢查並取用介面背後的實質型別。
空介面 interface{} 可以接受任何型別的值，但當實際處理資料時，必須知道空介面底下的實質型別為何。
例如、實質型別為字串或整數，就需要針對不同型別做不同的事。

在處理來源未知的 JSON 格式資料也是，Go 語言會用 map[string]interface{} 的型別來儲存解讀 JSON 字串後得到的結果；事先不知道解析出來的 JSON 資料會是數字或字串等等。
*/

// 當試圖直接轉換空介面的型別會？下面試試用 strconv.Atoi() 把底層值為字串的空介面轉為整數：
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s interface{} = "34"
	fmt.Println(strconv.Atoi(s))
}

//錯誤訊息指出 strconv.Atoi() 的參數不接受介面型別，必須做型別斷言才行。

/*
＊＊＊型別斷言語法：
v:=s.(T)

以上解釋為：用斷言「主張」介面值 s 底下的型別是 T，如果確實就將 T 型別的值 s 賦予 v。
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s interface{} = "34"
	v := s.(string)
	fmt.Println(strconv.Atoi(v))
}


// 當斷言轉換失敗時：
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i interface{} = 42
	s := i.(string)
	fmt.Println(strconv.Atoi(s))
}

// 斷言會傳回第二個值（布林值），指出轉換是否成功：
func main() {
	var str interface{} = 42
	if s, ok := str.(string); ok {
		fmt.Println(strconv.Atoi(s))
	} else {
		fmt.Println("Type assertion failed")
	}
}
