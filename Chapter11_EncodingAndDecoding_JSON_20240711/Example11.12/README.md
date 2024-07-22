# 11-6 gob：Go 自有的編碼格式 
* 以純文字為基礎的格式在解讀上仍然慢了點。例、JSON、XML。   
* 若系統完全以 Go 語言撰寫，那可以改用 Go 語言自己的二進位編碼格式—— gob。  
* Go 語言設計 gob 時，是以高效率、簡單易用和完整為考量，不需額外設定，甚至收發雙方使用的 Go 結構也不見得需要相同。  
* Go 1.16 下 gob 解碼及編碼的所需時間只有 JSON 的三分之一。  

gob 套件用起來和 json 的 Encoder/Decoder 很像：
```go
func NewEncoder(w io.Writer) *Encoder
func NewDecoder(r io.Reader) *Decoder
func (enc *Encoder) Encode(e interface{}) error
func (dev *Decoder) Decode(e interface{}) error
```  

練習、使用 gob 編碼和解碼資料
```go
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
```
顯示結果：
```
6b7f0301010773747564656e7401ff80000106010953747564656e74496401040001084c6173744e616d65010c00010d4d6964646c65496e697469616c010c00010946697273744e616d65010c00010a4973456e726f6c6c65640102000107436f757273657301ff840000001cff830201010d5b5d6d61696e2e636f7572736501ff840001ff82000032ff8103010106636f7572736501ff8200010301044e616d65010c0001064e756d6265720104000105486f75727301040000003fff80010401024675020453616c61010101030107456e676c69736801fe01960104000107486973746f727901fe026001040001044d61746801ffcc01060000
{2 Fu  Sala true [{English 203 2} {History 304 2} {Math 102 3}]}
```
