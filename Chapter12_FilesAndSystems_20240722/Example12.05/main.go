package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// 自訂 error
var ErrWorkingFileNotFound = errors.New("查無工作檔案")

func main() {
	workFileName := "note.txt"
	backupFileName := "backup.txt"
	err := writeBackup(workFileName, backupFileName)
	if err != nil {
		panic(err)
	}
}

// 備份檔案的函式
func writeBackup(work, backup string) error {
	workFile, err := os.Open(work) // 開啟工作檔
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrWorkingFileNotFound // 查無工作檔，傳回自訂 error
		}
		return err
	}
	defer workFile.Close() // 在備份結束後關閉工作檔

	backFile, err := os.OpenFile(backup, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) // 開啟備份檔，沒有就建立一個，資料附加到結尾
	if err != nil {
		return err
	}
	defer backFile.Close() // 在備份結束後關閉備份檔

	content, err := io.ReadAll(workFile) // 讀取工作檔內容
	if err != nil {
		return err
	}

	// 把一行日期和工作檔內容寫入備份檔
	backFile.WriteString(fmt.Sprintf("[%v]\n%v", time.Now().String(), string(content)))
	if err != nil {
		return err
	}
	return nil
}
