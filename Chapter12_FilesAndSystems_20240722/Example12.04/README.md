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
