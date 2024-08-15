# 13-6 查詢資料
資料表的查詢分兩類：  
1. 沒有參數、從資料表中取出大量資料用。
2. 有篩選條件，通常用來找出特定一筆符合的資料。


## 13-6-1 查詢並印出整個資料表內容
例、在資料表 employee 新增多筆資料，查詢整個內容並逐次印出每一筆資料各欄位的值：
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

	// 插入多筆資料
	insertStmt, err := db.Prepare("INSERT INTO employee(id,name) VALUE (?,?),(?,?)")
	if err != nil {
		panic(err)
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(306, "Pao", 307, "Ruby")
	if err != nil {
		panic(err)
	}
	fmt.Println("成功插入資料")

	// 查詢資料表，傳回 sql.Rows
	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println("資料表查詢成功")

	for rows.Next() { // 走訪 rows
		e := employee{}
		err := rows.Scan(&e.id, &e.name) // 讀出一筆資料
		if err != nil {
			panic(err)
		}
		fmt.Println(e.id, e.name) // 印出資料
	}
	err = rows.Err() // 檢查 Rows 有無遭遇其他錯誤
	if err != nil {
		panic(err)
	}
}
```
顯示結果：
```shell
sql.DB 結構已建立
資料庫連線成功
成功插入資料
資料表查詢成功
305 Olg
306 Pao
307 Ruby
```
```sql
MariaDB [mysqldb]> select * from employee;
+-----+------+
| id  | name |
+-----+------+
| 305 | Olg  |
| 306 | Pao  |
| 307 | Ruby |
+-----+------+
3 rows in set (0.001 sec)
```
解析：
> 1. 用 db.QUery() 來執行 SQL 指令。
> db.Exec() 和 Query() 差別，Query 會傳回 sql.Rows 結構，用來代表查詢結果的一列列資料，如下：
`func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error)`
> 2. 用 for 迴圈 rows.Next() 走訪它，迴圈每執行一次 rows 就會指向下列。此時可用 rows.Scan() 來將該列的欄位賦值給變數（變數數量必須跟欄位相同）
