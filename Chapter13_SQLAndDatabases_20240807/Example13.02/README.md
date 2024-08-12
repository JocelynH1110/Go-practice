# 13-3 以 Go 語言連接資料庫
經歷以上步驟後，連接資料庫其實是整個過程中最容易的一件事。  

但為了連接任何資料庫，必須先準備四件事：
1. 可供連接的資料庫伺服器
2. 使用者帳號
3. 使用者密碼
4. 特定操作的權限

資料庫權限包括：查詢、插入、移除資料、建立或刪除表格等等。

當完成前一節的準備，就可以開始撰寫 Go 程式了。  
1. 匯入相關套件：
```go
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
```
> "database/sql"：即 Go 內建的資料庫 API。
> _ "github.com/go-sql-driver/mysql"：匯入我們下載的 MySQL 驅動程式。
> _ 底線：意思是讓該套件使用底線為別名，因為不會直接使用 mysql 套件，只是要匯入它而已。（若匯入的套件有包括名稱，卻未呼叫他的功能，會在編譯時產生錯誤。

2. 連接資料庫：
* 看驅動程式和資料來源名稱格式是否有誤：
```go
	db, err := sql.Open("mysql", "jocelyn:1234@tcp(localhost:3306)/mysqldb?charset=utf8")
	if err != nil {
		panic(err)
	}
	fmt.Println("DB 結構已建立")
```
> 以上函式是 database/sql 套件提供的通用 API。
> 第一個參數：驅動程式名稱。
> 第二個參數：資料來源名稱（data source name）

第二個參數內容分別是：
`jocelyn:1234@tcp(localhost:3306)/mysqldb?charset=utf8` 
`使用者名稱:密碼＠tcp協定(伺服器位址)/資料庫名稱?UTF-8編碼`

sql.Open() 不會真正連線到資料庫，而是傳回一個 sql.DB 結構給我們使用。  
它傳回的 error，可用來檢查我們提供的驅動程式及資料來源名稱格式是否有誤。  


* 檢查資料是否可正確連線：
sql.Open() 並不知道所提供的帳號和密碼是否正確。  
若程式會長時間運作，也有可能遇到資料庫伺服器斷線或網路不穩問題，因此在任何操作前都應該先檢查資料庫的可連線狀況：
```go
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("資料庫連線成功")
```
> 在用完資料庫後，可以關閉資料庫連線。
> 正常情況下不需要這麼做，sql.DB 會在 Go 程式結束後自動關閉所有連線。
> 若資料庫會同時被數千人存取，而你對資料庫的操作不算頻繁、也只限於一個函數的範圍，那就該在函式結束時關閉資料庫連線。

* 關閉資料庫連線：
```go
defer db.Close()
```
