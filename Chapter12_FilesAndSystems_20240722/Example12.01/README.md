# 12-1 前言
* 本章將理解如何操作檔案系統，包括在磁碟上建立與修改檔案、及檢查檔案是否存在。  
* 也會實作一個命令列應用程式，可以接收各種旗標（flag）及其引數、以便控制你的程式行為、顯示說明文件等。
* 還會學到怎麼攔截作業系統發出的中斷訊息，並決定要在關閉程式之前做什麼處理。  

    * 命令列旗標與其引數
    * 系統中斷訊號
    * 檔案存取權限

# 12-2 命令列旗標與其引數
程式和系統的互動也不拘限於檔案本身，我們的程式也可以接收來自使用者的命令列旗標（flag），以便指定程式要做什麼。  

### 使用旗標
就像在終端機、命令提示字元或 Powershell 使用過一些命令列工具：
```go
go build -o ./bin/hello_world.exe main.go
```
> 這指令使用 go build 來將 main.go 編譯成 \bin 子目錄下的 hello_world.exe 可執行檔。
> 而執行檔的路徑與名稱就是透過旗標 -o （output）來指定。  
> 旗標有個名稱和對應值，而且可以是選擇性的、傳入順序也不必固定。  


對於旗標和引數，Go 語言提供了 flag 套件來協助開發者處理他們。  

* flag 套件提供了多種可定義旗標的函式，以下為常用的 flag：
```go
func String(name string, value string, usage string) *string

func Bool(name string, value bool, usage string) *bool

func Int(name string, value int, usage string) *int

func Int64(name string, value int64, usage string) *int64

func Float64(name string, value float64, usage string) *float64

func Duration(name string, value time.Duration, usage string) *time.Duration    // 時間長度

func Uint(name string, value uint,usage string) *uint   // uint 正整數

func Uint64(name string, value uint64,usage string) *uint64   // uint64 正整數

//分析所有命令行參數。
func Parse()

//Args 返回所有非 flag 的命令行參數切片，而 Arg 返回給定索引的非 flag 命令行參數。
func Args() []string

func Arg(i int) string
```
> 從以上函式的名稱與傳回值便可看出，其用途在接收特定型別的旗標引數。
> 每個函式都有以下三個參數：
    * name ：旗標的名稱，型別為字串。
    * value：旗標的預設值。
    * usage：說明旗標的用途（也是字串）。通常在設定旗標值錯誤時，這內容就會顯示給使者。

例子、旗標的使用
```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定義一個旗標 -value，接收整數，預設值為 -1
	v := flag.Int("value", -1, "Needs a value for the flag.")
	flag.Parse()
	fmt.Println(*v)
}
```
顯示結果：
```go
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722$ cd Example12.01/
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.01$ go run .
-1
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.01$ go run . -value 10
10
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.01$ go run . --value  30
30
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.01$ go run . --value=60
60
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.01$ go run . -value=80
80
```
> **解析：flag.Parse() 會解析使用者在命令列輸入的旗標 -value，並將其值以指標整數的形式賦予給 v。若沒有 -value 這個旗標名稱，*v 的值就是預設值 -1。**  

將以上檔案編成執行檔，再用 flag ：
```go
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.01$ go build -o main.exe 
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.01$ ./main.exe -value 100
100
```  

如果執行程式時加上 -h ，或者不確定旗標型別，程式會列出可用旗標、旗標值型別或其說明，然後結束：
```go
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.01$ ./main.exe -h
Usage of ./main.exe:
  -value int
    	Needs a value for the flag. (default -1)
```  

如果輸入錯誤型別：
```go
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.01$ ./main.exe -9b
flag provided but not defined: -9b
Usage of ./main.exe:
  -value int
    	Needs a value for the flag. (default -1)
```


### 以旗標來決定程式執行狀態
有時我們會希望某些旗標是執行應用程式的必要參數，查無此旗標的話就得提醒使用者。這代表得要謹慎決定旗標的預設值，因為得用這個預設值來判斷使用者是否有加上該旗標或給予正確值。

例、以下是更複雜的例子，程式最多只能接收三個旗標：
```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	n := flag.String("name", "", "Your first name")
	i := flag.Int("age", -1, "Your age")
	b := flag.Bool("married", false, "Are you married?")
	flag.Parse()

	if *n == "" { // 若名字旗標值為空字串，代表使用者沒有加上該旗標，或未給值
		fmt.Println("Name is required!")
		flag.PrintDefaults() // 印出所有旗標的預設值
		os.Exit(1)           // 結束程式
	}
	fmt.Println("Name：", *n)
	fmt.Println("Age ：", *i)
	fmt.Println("Married：", *b)
}
```
顯示結果：
```go
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.01$ go run .
Name is required!
  -age int
    	Your age (default -1)
  -married
    	Are you married?
  -name string
    	Your first name
exit status 1
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.01$ go run . -name jo -age 20 -married 
Name： jo
Age ： 20
Married： true
jocelyn@xps15:~/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.01$ go run . -name jo -age 20
Name： jo
Age ： 20
Married： false
```
> 當執行範例時，布林值指標後沒有引數，它仍收到 true 的值？這表示你可以用 flag.Bool() 定義一個不需引數的旗標，讓旗標本身當一個「開關」 。
> 若你沒加上 -married，效果就等同於寫 -married=false。
