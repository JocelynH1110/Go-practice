package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct { // 用來紀錄 employee 一筆資料的結構
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
	rowStmt, err := db.Prepare("SELECT name FROM employee WHERE id=?")
	if err != nil {
		panic(err)
	}
	defer rowStmt.Close()

	// 用參數化查詢來取出符合的單一一筆資料
	e := employee{id: 30}
	err = rowStmt.QueryRow(e.id).Scan(&e.name)
	if err != nil {
		panic(err)
	}
	fmt.Printf("id = %v 的員工名稱為 %v\n", e.id, e.name)
}
