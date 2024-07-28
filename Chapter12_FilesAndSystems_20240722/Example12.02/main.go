package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 建立訊號通道（緩衝區大小 1）
	sigs := make(chan os.Signal, 1)

	// 註冊要透過通道接收的訊號
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer cleanUp() // 延後執行的清理作業
	fmt.Println("程式執行中 （按下 ctrl+c 來中斷）")

Mainloop: // 一個標籤，用來代表以下這個無窮 for 迴圈
	for {
		s := <-sigs // 試著從通道接一個值
		switch s {  // 判斷收到的值是否為中斷或終止訊號
		case syscall.SIGINT:
			fmt.Println("程序中斷：", s)
			break Mainloop
		case syscall.SIGTERM:
			fmt.Println("程序中止：", s)
			break Mainloop
		}
	}
	fmt.Println("程式結束")
}

// 模擬程式中止後的清理作業
func cleanUp() {
	fmt.Println("進行清理作業...")
	for i := 0; i <= 10; i++ {
		fmt.Printf("刪除檔案 %v...(僅模擬)\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}
