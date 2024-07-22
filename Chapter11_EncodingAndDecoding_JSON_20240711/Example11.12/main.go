package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type student struct {
	StudentId     int
	LastName      string
	MiddleInitial string
	FirstName     string
	IsEnrolled    bool
	Courses       []course
}

type course struct {
	Name   string
	Number int
	Hours  int
}

func main() {
	// 原始資料
	s := student{
		StudentId:  2,
		LastName:   "Fu",
		FirstName:  "Sala",
		IsEnrolled: true,
		Courses: []course{
			{Name: "English", Number: 203, Hours: 2},
			{Name: "History", Number: 304, Hours: 2},
			{Name: "Math", Number: 102, Hours: 3},
		},
	}

	var conn bytes.Buffer            // 模擬通訊用的 io.Reader/io.Writer
	encoder := gob.NewEncoder(&conn) // 產生 encoder
	if err := encoder.Encode(&s); err != nil {
		fmt.Println("GOB 編碼錯誤：", err)
		os.Exit(1)
	}

	fmt.Printf("%x\n", conn.String()) // 把 conn 的內容用 16 進位形式印出

	s2 := student{}                             // 接收解碼後資料的結構
	decoder := gob.NewDecoder(&conn)            // 產生 decoder
	if err := decoder.Decode(&s2); err != nil { // 解碼 gob
		fmt.Println("GOB 解碼錯誤：", err)
		os.Exit(1)
	}
	fmt.Println(s2) // 解碼後的資料
}
