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

## 10-5-3 測量時間長度
時間長度的用途，不只能用來改變時間值。在現實中的應用程式裡，也可能會需要計算程式執行所耗費的時間。

#### 若要測量某段程式執行的時間，只需該段程式的頭尾各取一次當下系統時間，然後用第二個時間值的 Sub() 方法減去第一個時間值：
```
func (t Time) Sub(u Time) Duration
```

#### time.Since() 與 Until()
* 假如要用當下系統時間判斷某個時間到現在的時間長度，也可使用 time.Since(<時間值>)，這功能相當於以下寫法：
```
time.Now().Sub(<時間值>)
```
* 計算當下系統時間到未來某個時間還有多久：
```
time.Until(<時間值>)    // 相當於寫成 <時間值>.Sub(time.Now())
```

#### 練習、測試程式執行時間
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second * 2)
	end := time.Now()
	duration1 := end.Sub(start)    // 計算兩個時間值之間的長度
	duration2 := time.Since(start) // 計算 start 到 time.Now() 的時間長度

	fmt.Println("Duration1:", duration1)
	fmt.Println("Duration2:", duration2)

	// 檢查 duration1 是否小於 2500 毫秒（2.5秒）
	if duration1 < time.Duration(time.Millisecond*2500) {
		fmt.Println("程式執行時間符合預期")
	} else {
		fmt.Println("程式執行時間超出預期")
	}
}
```
顯示結果：
```
Duration1: 2.00112987s
Duration2: 2.001130293s
程式執行時間符合預期
```
