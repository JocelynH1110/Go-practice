# 12-3 系統中斷訊號
本章中，訊號（signal）指的是作業系統傳給我們的程式或程序的非同步通知。  
當程式收到訊號時，它會停下手邊的任務並設法處理這個訊號，可以的話就忽略它。  

在多數作業系統中，可以使用系統中斷訊號（signal）來對進行中的程式進行非標準的中斷，常見的情況是：  

當使用者在主控台按下 ctrl+c（^C）時，
* 系統會傳送名為 SIGINT 的中斷（interrupt）訊號給程式。
* 或者作業系統要強制終止程式，會傳送 SIGTERM 訊號給它。  
程式收到這些訊號時會立即結束，以 Go 程式來說就是執行 os.Exit(1)。

這樣的問題在於，就算程式內有使用 defer 延遲執行的程式（lesson 5），他們也不會被執行。且這些延遲執行的函式有可能負責以下的善後功能：
    * 釋出資源
    * 關閉資源
    * 結束資料庫連線  

可能會導致必要檢查無法完成。  

因此，我們可以在程式中註冊這些訊號，在收到訊號時能井然有序的完成該有的善後工作，並確保程式正常結束。  


### 接收中斷訊號通知
1. 若希望程式能判斷它何時收到特定的作業系統訊號，得使用 signal 套件的 Notify() 函式來註冊之：
```go
signal.Notify(<通知通道>, <訊號 1>, <訊號 2>...)
```
> 通知通道：當註冊的訊號發生時，它會被傳入通知通道。通道（channel）是 Go 語言中專門用於非同步程式資料交換的管道。（lesson 16）  
> 訊號：想接收的系統訊號，都定義在 syscall 套件的常數中。如、syscall.SIGINT（中斷）、syscall.SIGTERM（終止）等。  

2. 為了能註冊和收到系統訊號，必須先建立一個通道和用 make() 初始化它，以便能傳給 signal.Notify() 使用：
```go
<通道> := make(chan os.Signal, 1)
```
> chan（channel）：代表我們要建立通道，其內容型別為 os.Signal（及系統訊號）。
> 1：代表通道的緩衝區（buffer）大小為 1，也就是最多可暫存 1 個訊號。若想接收的訊號類型比較多，也可以加大緩衝區，緩衝區最少必須為 1，否則程式嘗試讀取通道時就很容易卡住。（lesson 16）  

3. 建立好通道後，可用以下方式取一個值出來：  
```go
<值> := <-<通道>
```
> 箭頭<-：受理算符（receive operators），意義是從通道取出一個值。
> 將值賦予給一個變數，以便拿來判斷內容。

練習、接收中斷訊號並優雅的結束程式：  
以下我們將攔截最常見的兩種訊號：syscall.SIGINT、syscall.SIGTERM。  

```go
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
```
顯示結果：
```go
程式執行中 （按下 ctrl+c 來中斷）
^C程序中斷： interrupt
程式結束
進行清理作業...
刪除檔案 0...(僅模擬)
刪除檔案 1...(僅模擬)
刪除檔案 2...(僅模擬)
刪除檔案 3...(僅模擬)
刪除檔案 4...(僅模擬)
刪除檔案 5...(僅模擬)
刪除檔案 6...(僅模擬)
刪除檔案 7...(僅模擬)
刪除檔案 8...(僅模擬)
刪除檔案 9...(僅模擬)
刪除檔案 10...(僅模擬)
```
> 為了能從 switch 內部直接脫離 for 迴圈，for 迴圈本身加上一個識別標籤叫 MainLoop。使得 break MainLoop 會打斷 for 迴圈，而不是單純的脫離 switch 敘述而已。
> 正常來說，應該將判斷訊號的程式碼放在一個非同步程序中，以免卡住其餘部份的程式。（lesson 16 會介紹建立非同步程序）

#### * syscall.SIGTERM 訊號
若你使用的是 Unix 系統，SIGTERM 訊號會在強行終止程式時傳送。  
如何觸發 syscall.SIGTERM：
1. 可用 go build 將練習題編譯成可執行檔，在一個終端機執行它：（沒有 exe 檔似乎也可以終止）
```go
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.02$ go build -o main.exe main.go
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.02$ ./main.exe
程式執行中 （按下 ctrl+c 來中斷）
程序中止： terminated
程式結束
進行清理作業...
刪除檔案 0...(僅模擬)
刪除檔案 1...(僅模擬)
刪除檔案 2...(僅模擬)
刪除檔案 3...(僅模擬)
刪除檔案 4...(僅模擬)
刪除檔案 5...(僅模擬)
刪除檔案 6...(僅模擬)
刪除檔案 7...(僅模擬)
刪除檔案 8...(僅模擬)
刪除檔案 9...(僅模擬)
刪除檔案 10...(僅模擬)
```
在開啟另一個終端機用 sudo ps -a 檢視所有程序 id，再以 sudo kill -<id> 終止它：
```
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.02$ sudo ps -a
    PID TTY          TIME CMD
  10159 pts/0    00:00:00 nvim
  11926 pts/1    00:00:00 go
  12069 pts/1    00:00:00 Example12.02
  12079 pts/2    00:00:00 sudo
  12081 pts/3    00:00:00 ps
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.02$ sudo kill 12069

// 沒有 sudo 似乎也沒差
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.02$  ps -a
    PID TTY          TIME CMD
  10159 pts/0    00:00:00 nvim
  12901 pts/1    00:00:00 go
  13053 pts/1    00:00:00 Example12.02
  13065 pts/2    00:00:00 ps
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.02$ kill 13053
```
