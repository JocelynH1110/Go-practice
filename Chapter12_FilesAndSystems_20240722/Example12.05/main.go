package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// 建立或開啟一個日誌檔
	// 其名稱為 log-年-月-日.txt，以當下時間為準
	logFile, err := os.OpenFile(fmt.Sprintf("log-%v.txt", time.Now().Format("2006-01-02")), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	// 建立 logger，寫入對象為前面開啟的檔案
	logger := log.New(logFile, "log:", log.Ldate|log.Lmicroseconds|log.Llongfile)

	// 將日誌輸出到檔案
	logger.Println("log message")
}
