# 12-5 建立與寫入檔案
## 12-5-1 用 os 套件新建檔案

os 套件的 Create() 方法能新建一個空白的新檔案，並賦予權限 0666 （所有使用者/群組都可讀/寫）。  
如果檔案已經存在，那該檔案的內容會被清空。
```go
func Create(name string)(*File, error)
```
> 成功新建或清空檔案後，os.Create() 會傳回一個 *os.File 結構。
> os.File 結構實作了 io.Reader 介面：也同時實作了 io.Writer 介面。  

例、以下程式會於程式目錄建立一個叫 text.txt 的文字檔，並在程式結束時以 File 結構的 Close() 關閉它：
```go
package main

import "os"

func main() {
	f, err := os.Create("text.txt") // 建立文字檔
	if err != nil {
		panic(err)
	}
	defer f.Close() // 確保在 main() 結束時關閉檔案
}
```


## 12-5-2 對檔案寫入字串
建立空檔案很簡單，但要對它寫入資料，檔案才會有內容。  

可以運用 os.File 的兩個方法：
```go
Write(b []byte)(n int, err error)
WriteString(s string)(n int, err error)
```
> Write() 和 WriteString() 功能是一樣的。只是接收的型別不同，一個是接收 []byte 切片、另一個是接收 string。
> 傳回值 n：代表函式對檔案寫入了 n 個位元，並會在寫入失敗時傳回非 nil 的 error，不過很多時候我們並不會接收這些值。  

例、新建檔案後，對該檔案結構寫入一些字串：
```go
package main

import "os"

func main() {
	f, err := os.Create("text.txt") // 建立文字檔
	if err != nil {
		panic(err)
	}
	defer f.Close() // 確保在 main() 結束時關閉檔案
	f.Write([]byte("使用 Write() 寫入\n"))
	f.WriteString("使用 WriteString() 寫入\n")
}
```
執行以上程式，同目錄下會出現 text.txt 其內容會如下：
```
使用 Write() 寫入
使用 WriteString() 寫入
```


## 12-5-3 一次完成建立檔案及寫入
Go 語言也允許用單一一個指令建立新檔案、並直接完成寫資料的動作。  

這要用到 os 套件的 WriteFile() 函式，定義如下：
```go
func WriteFile(filename string, data []byte, perm os.FileMode)error
```
> filename（字串）：檔案名稱。如果檔案不存在就會新建一個，而已經存在的檔案則會清空其內容。
> data []byte：要寫入的字串。
> perm：檔案權限。如前面介紹過的 0666、0763，這會用來設定新建檔案的權限。但若檔案已存在，就不會改變原有權限。


## 12-5-7 刪除檔案
可以使用 os.Remove() 函式：
```go
func Remove(name string)error
```
> 在刪除成功時候傳回值為 nil 的 error。

例、以下為一次完成建檔和寫入資料，及刪除檔案的例子：
```go
package main

import "os"

func main() {
	msg := "Hello Golang!"
	// 建立檔案並寫入資料
	err := os.WriteFile("text.txt", []byte(msg), 0644)
	if err != nil {
		panic(err)
	}

    // 刪除檔案
	rm := os.Remove("test.txt")
	if rm != nil {
		panic(rm)
	}
}
```


## 12-5-4 檢查檔案是否存在
上面的 os.Create()、os.WriteFile() 函式在碰上已經存在的檔案時，都會將其清空。  
Go 語言提供了檢查檔案存在與否的簡單機制。

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

// 檢查檔案是否存在的自訂函式
func checkFile(filename string) {
	finfo, err := os.Stat(filename) // 取得檔案描述資訊
	if err != nil {
		if errors.Is(err, os.ErrNotExist) { // 若 error 中包含檔案不存在錯誤
			fmt.Printf("%v:檔案不存在！\n\n", filename)
			return // 退出函式
		}
	}
	// 若檔案正確開啟，印出其檔案資訊
	fmt.Printf("檔名：%s\n是目錄：%t\n修改時間：%v\n權限：%v\n大小：%d\n\n", finfo.Name(), finfo.IsDir(), finfo.ModTime(), finfo.Mode(), finfo.Size())
}

func main() {
	checkFile("text.txt")
	checkFile("junk.txt")
}
```
結果顯示：
```go
檔名：text.txt
是目錄：false
修改時間：2024-07-30 16:46:54.376146602 +0800 CST
權限：-rw-r--r--
大小：13

junk.txt:檔案不存在！
```
> os.Stat() 方法傳回的錯誤可能包含多重 error 值，我們得檢查當中是否包含 os.ErrNotExist 錯誤，是的話就代表此檔案不存在。
> errors.Is(err, os.ErrNotExist) ：他的功能是檢查 err 是否是 os.ErrNotExist，或者說 err 否表示一個檔案或目錄不存在的錯誤。
> Go 1.13起擴充了錯誤檢查機制，官方建議使用 errors.Is(error, <欲檢查的錯誤值>) 來取代 os.IsNotExist() 等函式。
> Go 1.13 之前的版本中，得使用 os.IsNotExist(error) 來檢查 error 值是否包含 os.ErrNotExist 值。

以下為 Go 1.13 版本：
```go
package main

import (
	"fmt"
	"os"
)

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
```
顯示結果：
```go
<nil>
檔名：text.txt
是目錄：false
修改時間：2024-07-30 16:46:54.376146602 +0800 CST
權限：-rw-r--r--
大小：13
```


os.Stat() 及 os.File 結構的 Stat() 方法，會傳回一個 os.fileStat 結構，塌實做了 FileInfo 介面。  
這介面的方法能查詢檔案的各種資訊：
```go
type FileInfo interface{
    Name() string   // 檔名
    Size() int  // 檔案大小（計算方式取決於系統）
    Mode() FileMode // 修改權限
    ModTime() time.Time // 修改時間
    IsDir() bool    // 是否為目錄，相當等於呼叫 Mode().IsDir()
    Sys() interface{}   //檔案資料來源（有可能傳回 nil
}
```


## 12-5-5 一次讀取整個檔案內容 
在建立檔案後，自然會需要讀取它。若檔案不算太大的話，可以用本小節的兩個方式一口氣讀進所有內容。  
但若拿來開啟過大的檔案，就會耗掉大量系統的記憶體。下節會看如何一次只讀取一行字的做法。
### 使用 os.ReadFile()
第一種檔案全讀的方法如下：
```go
func ReadFile(filename string)([]byte, error)
```
> os.ReadFile() 會開啟檔名參數 filename 指定的檔案並讀取其內容，成功的話會以 []byte 切片形式傳回，error 會傳回 nil。
> os.File 結構在讀取內容時，或碰到檔案結尾會傳回 io.EOF（end of file）錯誤，但 ReadFile 是讀取整個檔案，故不會傳回 EOF。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("檔案內容：")
	fmt.Println(string(content))
}
```
顯示結果：
```
檔案內容：
Hello Golang!

```
### 使用 io.ReadAll() 搭配 os.Open()
第二種檔案全讀的方法：
```go
func ReadAll(r io.Reader)([]byte, error)
```
> 和第一種全讀功能很像，差別在於接收的參數是 io.Reader 介面型別。
> 這表示 ReadAll() 不只可以用來讀取 os.File 檔案，也能讀取符合 io.Reader 介面的任何物件，如 strings.NewReader() 或 http.Request 等。

若要讀取檔案，得先取得該檔案的 os.File 結構，辦法是使用 os.Open() 函式：
```go
func Open(name string)(*File, error)
```


```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("檔案內容：")
	fmt.Println(string(content))
}
```
