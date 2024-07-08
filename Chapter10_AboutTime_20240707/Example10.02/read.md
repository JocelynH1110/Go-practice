## 10-2-2 取得時間資料中的特定項目
例子、平時會做幾分鐘的簡易測試，但禮拜一凌晨0-2之間會執行全功能測試。程式需判斷現在時間是否允許進行全功能測試：
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	day := now.Weekday()
	hour := now.Hour()
	fmt.Println("Day:", day, "/hour:", hour)
	if day.String() == "Monday" && (hour >= 0 && hour < 2) {
		fmt.Println("執行全功能測試")
	} else {
		fmt.Println("執行簡易測試")
	}
}
```

### 轉換時間資料為字串
前一個範例中，時間資料的星期可以用 Weekday().String() 轉換字串，但不是 time.Time 的所有方法都能這樣做。Time 結構大多數方法傳回的資料就是 int 型別，而在 Go 語言中若想將 int 轉為 string，得使用 strconv 套件提供的轉換功能：

```go
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	appName := "HTTPCHECKER"
	action := "BASIC"
	date := time.Now()
	logFileName := appName + "_" + action + "_" + strconv.Itoa(date.Year()) + "_" + date.Month().String() + "_" + strconv.Itoa(date.Day()) + ".log"
	fmt.Println("log 檔名稱：", logFileName)
}
```
顯示結果：
```
log 檔名稱： HTTPCHECKER_BASIC_2024_July_8.log
```

* strconv.Itoa() 函式其實也會傳回 error 值，但既然時間資料的特定部份已知一定是 int 型別，就不必特地檢查轉換是否會失敗。
* 其實也可以用 fmt.Sprintf() 來產生格式化字串，這麼做也不需要自己轉換型別，只是轉換速度會比 strconv 套件慢。
logFileName := fmt.Sprintf("%v_%v_%v_%v_%v.log", appName, action, date.Year(), date.Month().String(), date.Day())
