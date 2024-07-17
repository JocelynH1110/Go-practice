package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	StudentId     int      `json:"id"`
	LastName      string   `json:"lname"`
	MiddleInitial string   `json:"minitial,omitempty"`
	FirstName     string   `json:"fname"`
	IsEnrolled    bool     `json:"enrolled"`
	Courses       []course `json:"classes,omitempty"`
}

type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

func main() {
	// 第一位學生資料，沒有課程資料，Courses 欄位為空值，故會被略過。
	s := student{
		StudentId:     1,
		LastName:      "Miller",
		MiddleInitial: "a",
		FirstName:     "Jay",
		IsEnrolled:    false,
	}

	// 編碼成 JSON ，縮排為 4 個空格
	student1, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(student1))
	fmt.Println()

	// 第二位學生資料，沒有中間名
	s2 := student{
		StudentId:  2,
		LastName:   "Julia",
		FirstName:  "Jane",
		IsEnrolled: true,
	}

	// 第二位學生選課資料
	c := course{Name: "English", Number: 102, Hours: 2}
	s2.Courses = append(s2.Courses, c)
	c = course{Name: "History", Number: 201, Hours: 2}
	s2.Courses = append(s2.Courses, c)
	c = course{Name: "Math", Number: 301, Hours: 3}
	s2.Courses = append(s2.Courses, c)

	student2, err := json.MarshalIndent(s2, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(student2))
}
