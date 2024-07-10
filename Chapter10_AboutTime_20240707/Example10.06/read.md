# 10-5 時間值的比較、時間長短處理
## 10-5-1 比較時間
有時候得確保 Go 程式必須在特定日期時間之前或之後執行特定任務。與其一個個比較時間值的各個部份，time 套件提供了更容易的方法來判斷兩個時間的先後順序。

#### 練習、比較時間順序
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(2050, 12, 31, 0, 0, 0, 0, time.Local)

	fmt.Println("Equal:", time.Now().Equal(date))   // 當下時間是否等於 date
	fmt.Println("Before:", time.Now().Before(date)) // 當下時間是否早於 date
	fmt.Println("After:", time.Now().After(date))   // 當下時間是否晚於 date
}
```
顯示結果：
```
Equal: false
Before: true
After: false
```

## 10-5-2 用時間長短來改變時間
AddDate() 只能用來更動日期（對時間值增減年、月、日）。
若要做出時、分、秒甚至小於 1 秒的改變，必須使用時間長度值（time.Duration 結構）來搭配時間值的 Add() 方法：
```
func (t Time) Add(d Duration) Time
```
* time.Duration 是自訂型別，代表時間的變量（單位為奈秒），或者說兩個時間值之間的差異，其底下的型別為 int6

#### 練習、使用 Duration 改變時間
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	duration1 := time.Duration(time.Second * 360)            // 時間長度 1（360 秒，等於 6 分鐘）
	duration2 := time.Duration(time.Hour*1 + time.Minute*30) // 時間長度 2（1 小時又 30 分鐘）
	fmt.Println("Dur1:", duration1.Nanoseconds(), "ns")      // 顯示時間長度值（以奈秒為單位）
	fmt.Println("Dur2:", duration2.Nanoseconds(), "ns")

	// 取得加上時間長度後的新時間
	date1 := now.Add(duration1)
	date2 := now.Add(duration2)

	// 顯示時間
	fmt.Println("Now:", now)
	fmt.Println("Date1:", date1)
	fmt.Println("Date2:", date2)
}
```
顯示結果：
```
Dur1: 360000000000 ns
Dur2: 5400000000000 ns
Now: 2024-07-10 17:19:07.305848325 +0800 CST m=+0.000008826
Date1: 2024-07-10 17:25:07.305848325 +0800 CST m=+360.000008826
Date2: 2024-07-10 18:49:07.305848325 +0800 CST m=+5400.000008826
```
