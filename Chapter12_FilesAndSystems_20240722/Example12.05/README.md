# 12-6 最完整的檔案開啟與建立功能：os.OpenFile()
前面已學習各種 **寫入檔案、建立\開啟檔案、讀取檔案** 的方法，這些以足以應付大多數狀況。  
若希望在開啟檔案時能指定更特定的行為，例如 **限制它採用唯讀或唯寫模式、要附加還是先清空其內容等**，就得用 os.OpenFile() 函式：
```go
func OpenFile(name string, flag int, pern FileMode)(*File, error)
```
> 成功開啟檔案時 OpenFile() 會傳回代表檔案的 os.File 結構。
> name：檔名。
> perm：要指定給新檔案的權限（八進位數），若要開啟的檔案不存在，且允許新建檔案時，新檔案就會套用這個權限。
> flag：它能決定檔案開啟後可進行哪些操作。

以下為 flag 參數在 os 套件中定義的一系列相關的常數：
```go
const(
    // 你必須指定 O_RDONLY, O_WRONLY 或 O_RDWR 其中之一
    O_RDONLY int = syscall.O_RDONLY     // 將檔案開啟為_唯讀模式
    O_WRONLY int = syscall.O_WRONLY     // 將檔案開啟為_唯寫模式
    O_RDWR int = syscall.O_RDWR         // 將檔案開啟為_可讀寫模式

    // 使用｜算符來連接以下旗標
    O_APPEND int = syscall.O_APPEND     // 將寫入資料附加到檔案尾端
    O_CREATE int = syscall.O_CREATE     // 檔案不存在時建立新檔案，如果已經存在，則沒有動作（寫入資料時會從頭覆蓋既有資料）
    O_EXCL int = syscall.O_EXCL         // 配合 O_CREATE 使用，確保檔案不存在
    O_SYNC int = syscall.O_SYNC         // I/O 同步模式（等待儲存裝置寫入完成）
    O_TRUNC int = syscall.O_TRUNC       // 開啟檔案時清空內容
)
```

#### 例、以上旗標可以串聯使用，改變檔案在不同情況下的操作行為：
```go
package main

import (
	"os"
	"time"
)

func main() {
	// 建立或開啟檔案
	f, err := os.OpenFile("junk.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write([]byte(time.Now().String() + "\n"))
}
```
> 若只是要開啟檔案（使用唯讀模式），使用 os.Open() 函式即可。


#### 練習、檔案備份  
把一個既存的文字檔 note.txt 的內容拷貝到備份檔 backupFile.txt，且寫入的內容還不能覆蓋既有資料，必須附加在檔案結尾才行。  

* 先寫入一個文字檔內容 note.txt：
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 建立或開啟檔案
	f, err := os.OpenFile("note.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	/* 第一種寫法
        for i := 1; i <= 10; i++ {
		fmt.Fprintf(f, "note %d\n", i)
	}*/

    // 第二種寫法
	for i := range 10 {
		fmt.Fprintf(f, "note %d\n", i+1)
	}
}
```
顯示結果：
```go
note 1
note 2
note 3
note 4
note 5
note 6
note 7
note 8
note 9
note 10
```

* 在使用 os.Open() 開啟 note.txt 時，假如檔案不存在，也得回傳一個自訂的 error 值：
```go
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
```
顯示結果：
```go
[2024-08-03 19:07:54.650136217 +0800 CST m=+0.000052034]
note 1
note 2
note 3
note 4
note 5
note 6
note 7
note 8
note 9
note 10
```
> 解析：上面練習題運用了前面提過的 os.Open()、io.ReadAll() 和檔案結構的 WriteString() 方法來備份來源工作檔案。  
為確保原有的備份資料不會被覆蓋，必須使用 os.OpenFile() 來指定他的寫入旗標之一為 os.O_APPEND。  

試著修改 note.txt 的內容，然後重複執行以上練習題。會發現程式將新版的工作檔內容複製到 backup.txt 的尾端，並還既護了備份的時間：
```go
[2024-08-03 19:07:54.650136217 +0800 CST m=+0.000052034]
note 1
note 2
note 3
note 4
note 5
note 6
note 7
note 8
note 9
note 10
[2024-08-03 19:18:16.469276036 +0800 CST m=+0.000031940]
note 0
note 1
note 2
note 3
note 4
note 5
note 6
note 7
note 8
note 9
note 10
```

#### 用 log 套件將日誌訊息寫入文字檔
將訊息輸出到檔案而不是主控台
```go
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
```
顯示結果：
```go
log:2024/08/07 13:50:40.691099 /home/jocelyn/working/Go-practice/Chapter12_FilesAndSystems_20240722/Example12.05/main.go:23: log message
```
