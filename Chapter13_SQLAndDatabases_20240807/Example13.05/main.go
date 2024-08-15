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
