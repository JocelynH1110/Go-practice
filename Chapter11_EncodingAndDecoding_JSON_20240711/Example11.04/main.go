package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	StudentId     int      `json:"id"`
	LastName      string   `json:"lname"`
	MiddleInitial string   `json:"minitial"`
	FirstName     string   `json:"fname"`
	IsEnrolled    bool     `json:"enrolled"`
	Courses       []course `json:"classes"`
}

type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

func main() {
	// JSON 資料
	data := []byte(`
	{
	"id":123,
	"lname":"Smith",
	"minitial":null,
	"fname":"John",
	"enrolled":true,
	"classes":[
		{
		"coursename":"Math",
		"coursenum":301,
		"coursehours":2
		},
		{
		"coursename":"History",
		"coursenum":302,
		"coursehours":3
		},
		{
		"coursename":"English",
		"coursenum":304,
		"coursehours":2
		}
	]
	}
	`)
	// 解析 JSON 並將值存入結構 s
	s := student{}
	// err := json.Unmarshal(data, &p)
	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", s) // 這樣可以印出欄位
	fmt.Println()
	fmt.Printf("%#v\n", s)
	fmt.Println()

	fmt.Println(s)
}
