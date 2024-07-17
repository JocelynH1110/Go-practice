## 11-3-4 有排版的 JSON 編碼結果
Marshal() 產生的 JSON 字串通通擠在一起，沒有縮排換行，當資料龐大時就不太容易閱讀。  

* 使用 MarshalIndent() 函式，作用跟 Marshal() 幾乎一樣，只差會縮排和換行：
```
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
```
> prefix 參數：是要放在每一行開頭的前綴詞。
> indent 參數：縮排文字，例如幾個空格或其他字元。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	LastName  string  `json:"lname"`
	FirstName string  `json:"fname"`
	Address   address `json:"address"`
}

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

func main() {
	// 建立要用來編碼 JSON 的資料結構
	addr := address{
		Street:  "Daan Road",
		City:    "Taipei",
		State:   "Taiwan",
		ZipCode: 106,
	}
	p := person{
		LastName:  "Miller",
		FirstName: "Kevin",
		Address:   addr, // 嵌入結構
	}

	// 編碼 JSON 資料但不排版
	noPrettyPrint, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(noPrettyPrint))
	fmt.Println()

	// 編碼 JSON 資料並排版
	PrettyPrint, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(PrettyPrint))
}
```
顯示結果：
```go
{"lname":"Miller","fname":"Kevin","address":{"street":"Daan Road","city":"Taipei","state":"Taiwan","ZipCode":106}}

{
	"lname": "Miller",
	"fname": "Kevin",
	"address": {
		"street": "Daan Road",
		"city": "Taipei",
		"state": "Taiwan",
		"ZipCode": 106
	}
}
```

練習、產生學生選課資料，然後轉成 JSON 格式傳給學生
```go
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
```
顯示結果：
```
{
    "id": 1,
    "lname": "Miller",
    "minitial": "a",
    "fname": "Jay",
    "enrolled": false
}

{
    "id": 2,
    "lname": "Julia",
    "fname": "Jane",
    "enrolled": true,
    "classes": [
        {
            "coursename": "English",
            "coursenum": 102,
            "coursehours": 2
        },
        {
            "coursename": "History",
            "coursenum": 201,
            "coursehours": 2
        },
        {
            "coursename": "Math",
            "coursenum": 301,
            "coursehours": 3
        }
    ]
}
```
