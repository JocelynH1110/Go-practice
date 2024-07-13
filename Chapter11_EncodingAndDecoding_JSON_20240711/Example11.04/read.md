## 11-2-3 解碼 JSON 到複合結構
### 以下為 JSON 資料，要用什麼樣的 Go 結構儲存？
```
	{
	"lname":"Smith",
	"fname":"John",
	"address":{
		"street":Daan Road",
		"city":"Taipei"
		"state":"Taiwan",
		"zipcode":800
    	}
	}
```
* 解析示範：
```go
package main

import (
	"encoding/json"
	"fmt"
)

type person struct { // 父結構
	Lastname  string  `json:"lname"`
	Firstname string  `json:"fname"`
	Address   address `json:"address"` // 子結構型別欄位
}

type address struct { // 子結構
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

func main() {
	// JSON 資料
	data := []byte(`
	{
	"lname":"Smith",
	"fname":"John",
	"address":{
		"street":"Daan Road",
		"city":"Taipei",
		"state":"Taiwan",
		"zipcode":106
		}
	}
	`)
	// 解析 JSON 並將值存入結構 p
	p := person{}
	// err := json.Unmarshal(data, &p)
	if err := json.Unmarshal(data, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", p) // 這樣可以印出欄位
}
```
顯示結果：
```
{Lastname:Smith Firstname:John Address:{Street:Daan Road City:Taipei State:Taiwan ZipCode:106}}
```  

### 練習、解碼學生課程 JSON 資料：
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
		}
		{
		"coursename":"English",
		"coursenum":304,
		"coursehours":2
		}
	]
	}

* 解析內容：
```go
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
	fmt.Printf("%+v\n", s)  // 這樣可以印出欄位
	fmt.Printf("%#v\n", s)  // 印出 main 和細節
	fmt.Println(s)          // 印出值
}
```

顯示結果：
```
{StudentId:123 LastName:Smith MiddleInitial: FirstName:John IsEnrolled:true Courses:[{Name:Math Number:301 Hours:2} {Name:History Number:302 Hours:3} {Name:English Number:304 Hours:2}]}

main.student{StudentId:123, LastName:"Smith", MiddleInitial:"", FirstName:"John", IsEnrolled:true, Courses:[]main.course{main.course{Name:"Math", Number:301, Hours:2}, main.course{Name:"History", Number:302, Hours:3}, main.course{Name:"English", Number:304, Hours:2}}}

{123 Smith  John true [{Math 301 2} {History 302 3} {English 304 2}]}
```
