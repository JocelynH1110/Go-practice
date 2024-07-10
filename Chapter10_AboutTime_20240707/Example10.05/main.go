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
