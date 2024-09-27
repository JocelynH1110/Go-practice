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
