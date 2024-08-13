# 13-4 建立、清空、移除資料表
確認資料庫連線正常後，要來替它建立資料表（table）。
建立資料表的目的，是要用一個抽象容器來存放彼此相關的資料，例如、員工出勤紀錄、統計數據等，不管是什麼資料，其共通的目的就是要讓應用程式讀取和解析之。  

SQL （Structured Query Language,結構化查詢語言）：一種用來操作 **關聯式資料庫** 系統的統一標準語言。
關聯式資料庫（relational database）：像是 MySQL、Postgres、DB2 都屬於這類資料庫，因此能用完全相同的方式操作他們。  

### 以下為資料表的 SQL 語法：
1. 建立資料表：
```sql
CREATE TABLE <資料表名稱>(
    <欄位 1 名稱> <資料型別> <限制>,
    <欄位 2 名稱> <資料型別> <限制>,
    <欄位 3 名稱> <資料型別> <限制>,
    ...
);
```
常見的資料型別和欄位的限制如下：
| 資料型別  | 欄位限制 |
| --- | --- |
|  INT (整數，也可以寫成 INTEGER) | NOT NULL (不得為 NULL 值)|
|  FLOAT (浮點數) | UNIQUE (必須是獨一無二的值)|
|  DOUBLE (雙精準度浮點數) | PRIMARY KEY (資料表主鍵)|
|  VARCHAR (字串，需指定長度) | FOREIGN KEY (資料表外鍵，即另一個資料表的主鍵)|

2. 清空資料表：
```sql
TRUNCATE TABLE <資料表名稱>;
```
3. 把資料表從資料庫中移除：
```sql
DROP TABLE <資料表名稱>;
```

例、在 MySQL 建立資料表：
| | |
| --- | --- |
| id | 員工代號|
| name | 名稱|
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
	fmt.Println("結構已建立")

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("資料庫連線成功")

    // 定義要使用的 SQL 指令
	DBCreate := `
	CREATE TABLE if not exists employee 
	(
		id INT NOT NULL UNIQUE,
		name VARCHAR(20)
	);`

    // 執行 SQL 指令
	_, err = db.Exec(DBCreate) // 執行 SQL 指令
	if err != nil {
		panic(err)
	}
	fmt.Println("表格 employee 已建立")
}
```
顯示結果：
```
結構已建立
資料庫連線成功
表格 employee 已建立
```
解析：
執行結果會指出 SQL 指令影響了資料多少筆資料，在此我們不需要知道，故用底線_跳過該值。  
只要傳回的 error 值為 nil，便代表資料表新增成功。  


### 在 MySQL 命令列客戶端 - 檢視表格資料表及移除
1. 檢視表格：
```sql
MariaDB [(none)]> use mysqldb;  
Database changed
MariaDB [mysqldb]> show tables;
+-------------------+
| Tables_in_mysqldb |
+-------------------+
| employee          |
+-------------------+
1 row in set (0.001 sec)
```

2. 移除表格：
```sql
MariaDB [mysqldb]> drop table employee;
Query OK, 0 rows affected (0.007 sec)
```
