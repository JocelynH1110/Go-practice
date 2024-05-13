// 7-4 在函式中活用介面
// 7-4-1 以介面為參數的函式
// io.Reader 介面可以用來接受不同型別的值。以下來看看其實際運用效果。

// 以下範例會寫出兩個任務相同的函式，用來解碼三筆 JSON 格式文字，但這兩個函是的參數型別不同，一個是字串，另一個是 io.Reader 介面。此外，前兩筆 JSON 資料是字串，但第三筆資料儲存在專案目錄下的文字檔 data.json，會被 Go 程式讀取成 io.File 檔案物件：

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Person struct { // 用於 JSON 資料的結構
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	s := `{"Name":"Jaja","Age":18}`
	s2 := `{"Name":"Haha","Age":28}`

	// 第一筆資料（字串）
	p, err := loadPerson(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)

	// 第二筆資料
	// strings.NewReader() 會傳回一個 strings.Reader 結構，符合 io.Reader 介面
	p2, err := loadPerson2(strings.NewReader(s2))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)

	// 第三筆資料
	// 讀取檔案後傳回 os.File 結構，符合 io.Reader 介面
	f, err := os.Open("data.json") // 開啟同資料夾下的文字檔
	if err != nil {
		fmt.Println(err)
	}
	p3, err := loadPerson2(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p3)
}

// 第一個 JSON 解析函式，接收字串參數
func loadPerson(s string) (Person, error) {
	var p Person
	err := json.NewDecoder(strings.NewReader(s)).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}

// 第二個 JSON 解析函式，接收 io.Reader 介面參數
func loadPerson2(r io.Reader) (Person, error) {
	var p Person
	err := json.NewDecoder(r).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}

/*
json 套件的 NewDecoder() 函式能解析 JSON 資料。
它實際上會接收一個 io.Reader 介面參數，並傳回解碼過的 Decoder 結構：

func NewDecoder(r io.Reader) *Decoder

然後程式會直接呼叫 Decoder 的 Decode() 方法，好將資料寫入 Person 結構變數的各個欄位。（lesson 11）

為了示範起見，函式 loadPerson() 會接收一個 string 型別引數，然後再呼叫 strings.NewReader() 把字串轉成 strings.Reader 結構傳給 json.NewDecoder() ；strings.Reader 即是一個實作了 io.Reader 介面的結構型別。至於在功能完全相同的函式 loadPerson2() 中，就直接接收一個 io.Reader 介面參數，重複完全一樣的過程。

io.Reader 介面定義：
type Reader interface{
	Read(p []byte) (n int,err error)
}

strings.Reader 和 os.File 定義： （他們都實作了上面的方法）
func (r *Reader) Read(b []byte) (n int,err error)

func (f *File) Read(b []byte) (n int,err error)


以上解釋了為何函式 json.NewDecoder() 能接收這些不同型別的值，並正確解讀出 JSON 資料。

當在開發 API 時，使用介面型別作為參數，就意味著使用者傳入的資料不會受限於特定型別、可用更大的彈性打造物件，只要符合介面的規範即可。
*/
