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
