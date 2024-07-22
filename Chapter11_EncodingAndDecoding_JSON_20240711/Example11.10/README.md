# 11-5 處理內容未知的 JSON 資料 
## 11-5-1 將 JSON 格式解碼成 map
有時候無法預知 JSON 的實際結果，比如某個網站 API 會產生動態的 JSON 回應，在不同情況下會有不同的鍵與值，或 API 本身規格會常常更新，導致傳回結果頻頻變動。  
json.Unmarshal() 不只能將 JSON 解碼，也能使用 map 來儲存資料。  

可將 map 定義成下面：

```go
map[string]interface{}      
```
> [string] 是 JSON 鍵
> interface{} 是 JSON 值  
> Unmarshal() 會將 JSON 資料中的任何鍵轉成 map 鍵，並將值配對給對應鍵。  
> JSON 鍵一定是字串，值可能是不同型別，所以用空介面接收。如此以來不管 JSON 資料中有什麼東西都可以放進 map。  

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := []byte(`{"checkNum":123,"amount":200,"category":["gift","clothing"]}`)

	// 定義 map
	var v map[string]interface{} // JSON 鍵一定是字串，值有可能是不同型別，故用空介面接收

	// 將 JSON 資料解碼到 map
	json.Unmarshal(jsonData, &v)

	// 印出 map 內容和走訪它
	fmt.Println(v)
	for key, value := range v {
		fmt.Println(key, "=", value)
	}
}
```
顯示結果：
```
map[amount:200 category:[gift clothing] checkNum:123]
amount = 200
category = [gift clothing]
checkNum = 123
```  
> 以上並未初始化以上並未初始化 map 變數 v，因為 Unmarshal() 會自己初始化。若 v 在傳入 Unmarshal() 前就有內容，那 Unmarshal() 會在 v 中新增其他鍵與值。  


練習、分析選課 JSON 資料內容：
```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonData := []byte(`
	{
		"id":2,
	"lname":"John",
	"fname":"Josh",
	"IsEnrolled":true,
	"grades":[100,76,93,50],
	"class":
		{
			"coursename":"History",
			"coursenum":101,
			"coursehours":3
		}
	}
	`)

	if !json.Valid(jsonData) { // 先檢查資料是否符合 JSON 格式
		fmt.Println("JSON 格式不合法：", jsonData)
		os.Exit(1)
	}

	// 解碼 JSON 格式到 map
	var v map[string]interface{}
	if err := json.Unmarshal(jsonData, &v); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 走訪 map
	for key, value := range v {
		fmt.Printf("%s = %v (%s)\n", key, value, findTypeName(value))
	}
}

// 用型別斷言來檢查值的函式
func findTypeName(i interface{}) string {
	switch i.(type) { // 型別斷言
	case string:
		return "string"
	case int:
		return "int"
	case float64:
		return "float64"
	case bool:
		return "bool"
	default:
		return fmt.Sprintf("%T", i)
	}
}
```
顯示結果：
```
grades = [100 76 93 50] ([]interface {})
class = map[coursehours:3 coursename:History coursenum:101] (map[string]interface {})
id = 2 (float64)
lname = John (string)
fname = Josh (string)
IsEnrolled = true (bool)
```
> 即使不曉得 JSON 資料中的鍵與值如何組成。仍能用一個 map[string]interface{} 容器來儲存解碼結果。
> 當每個值都透過空介面儲存，事後就得做型別斷言或使用 fmt.Sprintf() 才能判斷值的型別。
