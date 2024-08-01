package main

import (
	"fmt"
	"os"
)

// 檢查檔案是否存在的自訂函式
func main() {
	finfo, err := os.Stat("junk.txt")
	if err != nil {
		if os.IsNotExist(err) {
			//fmt.Printf("%v:檔案不存在！\n\n", finfo)
			fmt.Println(finfo)
		}
	}
	finfo, err = os.Stat("text.txt")
	if err != nil && os.IsNotExist(err) {
		fmt.Println("text")
	}
	fmt.Printf("檔名：%s\n是目錄：%t\n修改時間：%v\n權限：%v\n大小：%d\n\n", finfo.Name(), finfo.IsDir(), finfo.ModTime(), finfo.Mode(), finfo.Size())
}
