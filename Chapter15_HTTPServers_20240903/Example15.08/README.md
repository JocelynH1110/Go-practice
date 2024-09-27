## 15-5-3 使用模板檔案產生動態網頁
通常靜態資源會原封不動供人存取，若想傳回動態內容網頁，也可使用外部模板檔案。這能在不重啟伺服器的情況下修改模板。  

為提升效能，可選擇在伺服器啟動時就載入模板，並把它儲存在請求處理器結構中：若很重視網頁存取速度，考量到會有多重客戶端連線，便可考慮這麼做。但如此一來，若想修改模板，就仍得重開伺服器。（可使用並行性運算繞過這限制，Lesson 16）  

現在要把模板變成外部資源，在一個 HTML 檔案裡寫模板語言，然後從伺服器載入它。模板引擎會解讀之，把參數傳入填空處。  
對於此，可以使用 html/template 函式：
```go
func ParseFiles(filenames ...string)(*Template, error)
```

使用方式：
1. 呼叫它
```go
tmpl, err:=template.ParseFiles("mytemplate.html")
```

2. 再呼叫 `teml.Execute()` 就可產生內容。


### 使用外部模板檔案
1. 建立外部模板檔案，內容就是模板字串：
```html
<html>
    <h1>Customer {{.ID}}</h1>
    {{if .ID}}
        <p>Details:</p>
        <ul>
            {{if .Name}}<li>Name: {{.Name}}</li>{{end}}
            {{if .Surname}}<li>Surname: {{.Surname}}</li>{{end}}
            {{if .Age}}<li>Age: {{.Age}}</li>{{end}}
        </ul>
    {{else}}
        <p>Data not available</p>
    {{end}}
</html>
```

2. 請求處理器結構，並用欄位 tmpl 來紀錄模板檔案內容：
```go
package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Customer struct {
	ID      int
	Name    string
	Surname string
	Age     int
}

// 會紀錄模板的請求處理器
type Hello struct {
	tmpl *template.Template
}

// 請求處理器方法
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vl := r.URL.Query()
	customer := Customer{}

	id, ok := vl["id"]
	if ok {
		customer.ID, _ = strconv.Atoi(id[0])
	}

	name, ok := vl["name"]
	if ok {
		customer.Name = name[0]
	}

	surname, ok := vl["surname"]
	if ok {
		customer.Surname = surname[0]
	}

	age, ok := vl["age"]
	if ok {
		customer.Age, _ = strconv.Atoi(age[0])
	}

	h.tmpl.Execute(w, customer)
}

func main() {
	// 建立請求處理器
	hello := Hello{}

	// 載入模板檔案和建立模板物件，賦予給請求處理器
	hello.tmpl, _ = template.ParseFiles("./template/hello_tmpl.html")

	// 註冊請求處理器
	http.Handle("/", hello)

	// 啟動伺服器
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
