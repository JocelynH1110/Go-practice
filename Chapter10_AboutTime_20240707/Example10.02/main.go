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
