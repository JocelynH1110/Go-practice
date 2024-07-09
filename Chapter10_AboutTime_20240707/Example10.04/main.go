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
