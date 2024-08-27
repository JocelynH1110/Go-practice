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
