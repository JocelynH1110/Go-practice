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
