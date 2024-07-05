## 9-3-3 建立自訂 logger 物件
到目前為止，本章使用 log 印出日誌時，都是透過該套件提供的標準 logger（standard logger）。

也可以依需要建立自己的（多重）logger，以便針對不同的情境輸出訊息：
```
<logger> := log.New(<io.Writer 介面>, <前綴詞>, <旗標>)
```

* 標準 logger 的第一個參數會使用 os.Stdout，這符合 io.Writer 介面的物件其實就是將訊息印出到主控台。
* 也可以換成其他物件，os.File 結構就是另一個符合 io.Writer 介面的東西，這使你能將日誌訊息寫到檔案中。
* 前綴詞是個字串，會加在 log 訊息最前面，除非用 log.Lmsgprefix 旗標讓它挪到使用者自己的訊息之前。
* 旗標參數則和前面使用 SetFlags() 設定的一樣，能用來決定 logger 訊息的格式。

```go
package main

import (
	"errors"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "log:", log.Ldate|log.Lmicroseconds|log.Llongfile)
	logger.Println("Start of our app")
	err := errors.New("application aborted!")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("End of our app")
}
```
結果輸出：
```
log:2024/07/05 14:06:39.519233 /home/jocelyn/working/Go-practice/Chapter09_BasicDebugging_20240614/Example09.05/main.go:11: Start of our app
log:2024/07/05 14:06:39.519303 /home/jocelyn/working/Go-practice/Chapter09_BasicDebugging_20240614/Example09.05/main.go:14: application aborted!
exit status 1
```

