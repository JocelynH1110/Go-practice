## 8-4-2  init() 函式 
GO 語言中的套件分為兩種：可執行的、不可執行的。

main 套件是特殊套件，也是可執行的。
main套件裡一定要有一個 main() 函式，這是 Go 語言在 go run、go build 等指令會尋找並執行的對象。

任何套件檔案 —包括 main 套件在內— 還可定義一個特殊的函式 init()，它可用來替套件設置初始狀態或初始值。
* 以下是運用 init() 的例子：
1. 設置資料庫物件和連線
2. 初始化套件變數
3. 建立檔案
4. 載入設定組態
5. 驗證或修復程式狀態


* 對於一個套件檔案，Go 語言會以下面的順序呼叫 init() 和 main()：
1. 匯入的外部套件的套件層級最先初始化
2. 接著套件自身的套件層級會初始化
3. 呼叫外部套件的 init()
4. 呼叫套件本身的 init()
5. 如果執行的檔案是 main 套件，最後會呼叫套件本身的 mian() 套件


以下是示範 init() 和 main() 執行順序的例子：
```go
package main

import(
    "fmt"
)

var name = "Gordo"

func init(){
    fmt.Println("Hallo,",name)
}

func main(){
    fmt.Println("Halo,main() 函式")
}
```

執行結果：
`Hallo,Gordo`
`Hallo,main() 函式`

分析：
套件層級的宣告變數會最先執行，name -> init() -> main()

注意、 init() 函式不能有參數和回傳值，會產生錯誤。
