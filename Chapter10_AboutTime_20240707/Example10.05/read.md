## 10-4-2 設定時區來取得新時間值
### 除了在使用 New() 建立 Time 結構時設定時區，也可用 Time 結構本身的 In() 來指定時區，並傳回一個新時間值：
```
func (t time) In(loc *Location) Time
```

### 想使用特定的時區、甚至建立自訂時區，有以下兩種方法：
```
func LoadLocation(name string) (*Location, error)
func FixedZone(name string, offset int) *Location
```
* LoadLocation() 函式要傳入一個現有的 IANA 時區名稱，以此來建立時區結構（失敗時傳回不為 nil 的 error）。
* FixedZone() 以 UTC 時區為準，加減 offset 填入的秒數後，傳回一個以 name 參數為名稱的自訂地區。

#### 練習、設定不同時區
```go
package main

import (
	"fmt"
	"time"
)

func displayTimeZone(t time.Time) {
	fmt.Print("Time:", t, "\nTimezone:", t.Location(), "\n\n")
}
func main() {
	// 本地時間
	// date := time.Date(2021, 4, 22, 16, 44, 05, 324359102, time.Local)
	date := time.Now()
	// 設為美國紐約地區
	timeZone1, _ := time.LoadLocation("America/New_York")
	// 美國紐約時區
	remoteTime1 := date.In(timeZone1)
	// 設為澳洲雪梨時區
	timeZone2, _ := time.LoadLocation("Australia/Sydney")
	// 澳洲雪梨時區
	remoteTime2 := date.In(timeZone2)
	// 自訂時區
	timeZone3 := time.FixedZone("MyTimeZone", -1*60*60)
	// 自訂時區，即 UTC 時區減 1 小時
	remoteTime3 := date.In(timeZone3)

	displayTimeZone(date)
	displayTimeZone(remoteTime1)
	displayTimeZone(remoteTime2)
	displayTimeZone(remoteTime3)
}
```
顯示結果：
```
Time:2024-07-10 15:01:22.453426709 +0800 CST m=+0.000009438
Timezone:Local

Time:2024-07-10 03:01:22.453426709 -0400 EDT
Timezone:America/New_York

Time:2024-07-10 17:01:22.453426709 +1000 AEST
Timezone:Australia/Sydney

Time:2024-07-10 06:01:22.453426709 -0100 MyTimeZone
Timezone:MyTimeZone
```
