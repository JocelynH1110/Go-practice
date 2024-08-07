# 12-7 處理 CSV 格式檔案
除了純文字檔和 JSON 資料外，程式最常存取的檔案格式之一：CSV（comma-separated value, 逗號分隔值）。  
CSV 本身是純文字，但使用逗號來分隔每一直行（column）或欄位的值，而每一橫列（row）則以每一行結尾的換行符號區隔。  

以下為經典 CSV 格式資料：
```
firstName,lastName,age
Celia,Jones,30
Case,Jo,25
```
> 第一行是標頭（header），即各行或欄位的名稱。
> csv 的分隔符號是半形逗號，有時也可能用空格或 tab（\t）等字元。  


## 12-7-1 走訪 CSV 檔內容
Go 語言也提供了標準函式庫 encoding/csv，可用來解析 CSV 格式資料：
```go
func NewReader(r io.Reader) *Reader
```
>  bufio 套件一樣，csv 套件的 NewReader() 接收一個 io.Reader 介面型別，後傳回 csv.Reader 結構。  

此結構的 Read() 方法能用來讀取 csv 資料的一行內容，並將其轉換成 []string 切片：
```go
func (r *Reader) Read() (record []string, err error)
```
> 不同於 bufio.Reader 的 ReadString() 方法，csv 套件會自動判斷結尾的換行符號。但這行文字會以字串切片的形式傳回。  

#### 例、從一個名為 data.csv 的檔案讀取 CSV 資料，data.csv 檔案內容寫入，如下：
```go
package main

import "os"

func main() {
	f, err := os.OpenFile("data.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("firstName,lastName,age\nCelina,Joned,18\nCailyn,Hely,28\nCayden,Smith,42")
}
```

只要開啟上面這個 csv 檔案，就能將 os.File 結構傳給 csv.NewReader() 來建立所需的物件：
```go
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("data.csv") // 開啟 CSV 檔案
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file) // 取得 csv.Reader 結構
	for {
		record, err := reader.Read() // 從 csv.Reader 讀取一行資料
		if err == io.EOF {           // 遇到檔案結尾錯誤，就離開迴圈
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(record) // 印出該行資料
	}
}
```
顯示結果：
```go
[firstName lastName age]
[Celina Joned 18]
[Cailyn Hely 28]
[Cayden Smith 42]
```


## 12-7-2 讀取每行資料各欄位的值 
如前所提，csv.Reader 的 Read() 方法會傳回 []string 切片。 csv 套件會以半形逗號為依據，將各欄位轉成字串切片的不同元素：索引 0 是左邊數來第一個欄位，索引 1 是第二個...。  
所以只要事先知道 CSV 檔的欄位組成，就很容易取出想要的東西。
|  索引 0  |  索引 1  |  索引 2  |
| --- | --- | --- |
| firstName | lastName | age |
| Celina | Jones | 20 |

跳過標頭的話，可以加入一個布林變數做為開關，在 for 迴圈第一次執行時選擇不印出東西。

#### 例、修改上節範例，跳過標頭：
```go
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

const (
	firstName = iota // CSV 欄位索引
	lastName
	age
)

func main() {
	file, err := os.Open("data.csv") // 開啟 CSV 檔案
	if err != nil {
		panic(err)
	}
	defer file.Close()

	header := true                // 標頭開關
	reader := csv.NewReader(file) // 取得 csv.Reader 結構
	for {
		record, err := reader.Read() // 從 csv.Reader 讀取一行資料
		if err == io.EOF {           // 遇到檔案結尾錯誤，就離開迴圈
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if header {
			header = false
			continue // 跳過第一行（標頭）
		}
		fmt.Println("--------------")
		fmt.Println("First name:", record[firstName])
		fmt.Println("Last name:", record[lastName])
		fmt.Println("Age:", record[age])
	}
}
```
顯示結果：
```go
--------------
First name: Celina
Last name: Joned
Age: 18
--------------
First name: Cailyn
Last name: Hely
Age: 28
--------------
First name: Cayden
Last name: Smith
Age: 42
```
