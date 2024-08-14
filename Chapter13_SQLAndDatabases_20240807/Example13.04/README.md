# 13-5 插入資料
在 Go 語言中，對資料表插入新資料的動作分成兩個階段：
 1. 先用 sql.DB 結構的 Prepare() 產生 **參數化查詢敘述（prepared statement）** ，也就是個 sql.Stmt 結構。
 2. 再透過這個來實際操作資料庫。

以上會這樣做是為了避免 SQL 注入（SQL injection）。  

攻擊者可以故意輸入 <名稱> OR '1'='1' 或類似字串，使得網站用來查詢資料庫的 SQL 指令變如下：
```
"SELECT <密碼> FROM <帳號資料表> WHERE <帳號名稱> = <名稱> OR '1'='1';"
```
> 由於 '1'='1' 必然為真，使前面的名稱檢查條件因為 OR 算符而失去作用，結果就傳回資料表中所有的密碼。  

* 參數化查詢：目前公認防範 SQL 注入最有效的辦法，因為資料庫會先將 SQL 指令編譯成位元組碼，然後才透過參數將值放進需要的地方。  
也就是說，就算傳入參數的值帶有 SQL 指令，它也不會被資料庫執行。這樣做更能提高 SQL 執行效率，因為一部分的指令已經事先編譯好了。  


例、在 MySQL 資料表插入資料:
```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "jocelyn:1234@tcp(localhost:3306)/mysqldb?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("sql.DB 結構已建立 ")

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("資料庫連線成功")

	// 準備參數化查詢敘述
	insertStmt, err := db.Prepare("INSERT INTO employee(id,name) VALUES (?,?);")
	if err != nil {
		panic(err)
	}
	defer insertStmt.Close()             // 在程式結束時，關閉參數化查詢敘述
	_, err = insertStmt.Exec(305, "Olg") // 新增一筆資料
	if err != nil {
		panic(err)
	}
	fmt.Println("成功插入資料 305,Olg")
}
```
顯示結果：
```shell
sql.DB 結構已建立 
資料庫連線成功
成功插入資料 305,Olg
```
> 為避免 SQL 注入攻擊，要填入的值先寫成問號，代表參數，並由 db.Prepare() 編譯和傳回 insertStmt 結構。  
> insertStmt.Exec() 會將填入的值放進上述的 INSERT 指令中問號的所在位置。

不同資料庫的參數化查詢語法會有所不同。例、postgres 是使用 $1、$2 ，這樣的符號來代表參數。  

參數化查詢敘述（sql.Stmt 結構）會佔用一些資源，這取決於資料庫類型，有可能是資料庫後端或驅動程式本身的資源。  
故在用完 sql.Stmt 結構後，應該用 Close() 關閉他來釋放資源。
