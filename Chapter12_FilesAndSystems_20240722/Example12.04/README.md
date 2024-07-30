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
使用 WriteString() 寫入1
```
