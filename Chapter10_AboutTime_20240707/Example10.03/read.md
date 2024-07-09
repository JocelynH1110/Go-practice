# 10-3 時間值的格式化
## 10-3-1 將時間轉成指定格式的字串
time.Time 結構的 Format() 方法可以將時間轉成特定格式的字串：
```
func (t time) Format(layout string) string
```
參數 layout 為時間格式字串。

* 練習、用不同格式輸出時間字串（將 time.Now() 傳回的系統時間轉成不同的格式並印出）
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(time.ANSIC))       // 美國國家時間格式
	fmt.Println(time.Now().Format(time.UnixDate))    // Unix 系統格式
	fmt.Println(time.Now().Format(time.RFC3339))     // RFC3339 格式
	fmt.Println(time.Now().Format("2006/1/2 3:4:5")) // 自訂格式
}
```
顯示結果：
```
Mon Jul  8 15:12:56 2024
Mon Jul  8 15:12:56 CST 2024
2024-07-08T15:12:56+08:00
2024/7/8 3:12:56
```

## 10-3-2 將特定格式的時間字串轉成時間值
Go 語言也允許將符合特定格式的時間字串轉成 time.Time 結構。轉換成 Time 結構後就能照想要的方式格式化他了，輸出閱讀性好的多的結果。time 套件也提供這方面的功能：
```
func Parse(layout ,value string) (Time,error)
```
* Parse() 會試圖以 layout 參數指定的格式轉換 value 中的日期時間。
* 若格式不符而導致轉換失敗，那 Parse() 會傳回一個存有時間零值的 Time 結構，以及不為 nil 的 error 值。
* 時間零值指的是 January 1,year 1,00:00:00 UTC ，而不是結構本身的零值。要檢查 Time 結構時間是否為零值，可用 IsZero() 方式來判斷。

### 練習、將時間字串轉成 time.Time 結構

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t1, err := time.Parse(time.ANSIC, "Thu Apr 22 16:44:05 2021") // 美國國家時間格式
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("from ANSIC :", t1)

	t2, err := time.Parse(time.UnixDate, "Thu Apr 22 16:44:05 CST 2021") //Unix 系統格式
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("from UnixDate:", t2)

	t3, err := time.Parse(time.RFC3339, "2021-04-22T16:44:05+08:00") //RFC3339 系統格式
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("from RFC3339:", t3)

	t4, err := time.Parse("2006/1/2 PM 3:4:5", "2021/04/22 PM 4:44:5") // 自訂格式
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("from custom:", t4)
}
```
結果顯示：
```
from ANSIC : 2021-04-22 16:44:05 +0000 UTC
from UnixDate: 2021-04-22 16:44:05 +0800 CST
from RFC3339: 2021-04-22 16:44:05 +0800 CST
from RFC3339: 2021-04-22 16:44:05 +0000 UTC
```
Unix 系統格式時間與 RFC3339 格式會含有時區資訊（CST 時區，即 UTF+8），但 ANSIC 時間沒有，因此被 time 套件認定為 UTC 標準時間。
