## 9-3-2 用 log 套件輸出日誌
### 使用 log.Printin()
比起 fmt.Println() 或類似的功能輸出追蹤訊息，Go 語言的 log 套件能用更豐富的細節來紀錄程式執行資訊。

* log.Println 的使用：
```go
package main

import (
	"errors"
	"log"
)

func main() {
	log.Println("Start of our app")
	err := errors.New("application aborted!")
	if err != nil {
		log.Println(err)
	}
	log.Println("End of our app")
}
```

執行結果：
```
2024/07/04 23:40:20 Start of our app
2024/07/04 23:40:20 application aborted!
2024/07/04 23:40:20 End of our app
```
比起 fmt.Println() 多了訊息的時間搓記，對日後檢視日誌、設法釐清臭蟲的發生時間和順序非常有用。
log 套件的 Print()、Printf() 的使用方式和 fmt 套件的同名函式是一樣的。

* 自訂 log 套件的日誌格式（Setflags() 函式）：
```go
func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("Start of our app")
	err := errors.New("application aborted!")
	if err != nil {
		log.Println(err)
	}
	log.Println("End of our app")
}
```

執行結果：
```
2024/07/05 00:02:13.429680 /home/jocelyn/working/Go-practice/Chapter09_BasicDebugging_20240614/Example09.04/main.go:10: Start of our app
2024/07/05 00:02:13.429768 /home/jocelyn/working/Go-practice/Chapter09_BasicDebugging_20240614/Example09.04/main.go:13: application aborted!
2024/07/05 00:02:13.429778 /home/jocelyn/working/Go-practice/Chapter09_BasicDebugging_20240614/Example09.04/main.go:15: End of our app

```
log.SetFlags() 以聯集算符串聯了三個旗標，這些旗標都是 log 套件提供的常數。
此時日誌訊息時間會精確到微秒，並會包括原始檔的完整路徑。

* log.SetFlags 可用的旗標
```go
const(
    Ldate =1 << iota    //本地日期（年、月、日）
    Ltime               //本地時間（時、分、秒）
    Lmicrosecond        //本地時間（時、分、秒、微秒）
    Llongfile           //寫出完整路徑、檔名、程式行號
    Lshortfile          //寫出簡短檔名、程式行號
    LUTC                //若有使用 Ldate、Ltime 旗標，改顯示 UTC 時間
    Lmsgprefix          //若有前綴詞，將之挪到使用者自己的訊息前面
    LstdFlags =Ldate|Ltime  //logger 的預設旗標
)
```

* 正常情況下，前綴詞會放在整個 log 訊息的最前面：
```
<前綴詞> <時間> <程式名稱>: <使用者訊息>
```

* 若啟用 Lmsgprefix 旗標，前綴詞會挪到使用者訊息前面：
```
<時間> <程式名稱>: <前綴詞> <使用者訊息>
```
也可以用 log.SetPrefix() 或 log.New() 設定日誌訊息的前綴詞。


### 使用 log.Fatal()、log.Panic() 紀錄嚴重錯誤
* Fatal()、Fatalln()、Fatalf() 方法的作用跟 log 或 fmt 的 Print()、Println()、Printf() 相同。
* 差別在於：Fatal() 們在輸出訊息後，接著會呼叫 os.Exit(1) 來中止程式。
* Panic()、Panicln()、Panicf() 用法和 Fatal() 系列相同。
* 差別在於：會引發 panic。panic 可以用 recover() 函式救回來，但 os.Exit 就不行了。 （lesson 6）

當有重大錯誤發生時，可在輸出日誌追蹤資訊的同時決定是否要中止程式，該不該給使用者機會挽救。
若錯誤可能會令應用程式的資料受損、或發生難以預期的行為，那最好的是在事態惡化前先讓程式當掉。
若程式結束時需要做一些安全操作，如透過用 defer 延遲執行的函式來關閉檔案或資料庫，那使用 log.Panic() 會是叫好的選擇。

例子、讓程式在遇到錯誤時當掉：
```go
func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("Start of our app")
	err := errors.New("application aborted!")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("End of our app")
}
```
執行結果：
```
2024/07/05 13:46:39.559488 /home/jocelyn/working/Go-practice/Chapter09_BasicDebugging_20240614/Example09.04/main.go:10: Start of our app
2024/07/05 13:46:39.559535 /home/jocelyn/working/Go-practice/Chapter09_BasicDebugging_20240614/Example09.04/main.go:13: application aborted!
exit status 1
```
