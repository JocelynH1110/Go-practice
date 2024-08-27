# 13-7 更新既資料
更新既有資料跟插入資料的動作是很像的。

例、更新資料
```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "jocelyn:1234@tcp(localhost:3306)/mysqldb?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("sql.DB 結構已建立")
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("資料庫連線成功")

	// 產生參數化查詢敘述
	updateStmt, err := db.Prepare("UPDATE employee SET name=? WHERE id=?")
	if err != nil {
		panic(err)
	}
	defer updateStmt.Close()

	// 更改原有資料，並執行參數化查詢
	e := employee{305, "Mary"}
	updateResult, err := updateStmt.Exec(e.name, e.id)
	if err != nil {
		panic(err)
	}

	// 檢查更新時影響了幾筆資料
	updateRecords, err := updateResult.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("更新資料比數：", updateRecords)
}
```
顯示結果：
```shell
sql.DB 結構已建立
資料庫連線成功
更新資料比數： 0
```
解析：  
更新的過程跟插入資料幾乎一樣，都是先用 db.Prepare() 產生參數化查詢敘述 updateStmt 後，再用 updateStmt.Exec() 來執行和傳入參數。  
在此也用了 RowsAffected() 方法來檢視更新時影響了幾筆資料。
