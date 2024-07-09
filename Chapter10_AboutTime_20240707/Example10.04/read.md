# 10-4 時間值的管理
## 10-4-1 建立和增減時間值
* 建立代表特定時間的 time.Time 結構：
```
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
```
* 建立 Time 結構後，可以用他的 AddDate() 方法來增減其日期：
```
func (t time) AddDate(years int, month int, days int) Time
```
### 建立並改變時間值
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	date1 := time.Date(2021, 4, 22, 16, 44, 05, 324359102, time.UTC) // 使用 UTC 時區
	fmt.Println(date1)
	date2 := time.Date(2021, 4, 22, 16, 44, 05, 324359102, time.Local) // 使用本地時區
	fmt.Println(date2)
	date3 := date2.AddDate(-1, 3, 5) // 減 1 年，加 3 個月又 5 天
	fmt.Println(date3)
}
```
顯示結果：
```
2021-04-22 16:44:05.324359102 +0000 UTC
2021-04-22 16:44:05.324359102 +0800 CST
2020-07-27 16:44:05.324359102 +0800 CST
```
