package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// HTML 模板原始字串
var templateStr = `
<html>
	<h1>Customer{{.ID}}</h1>
	{{if .ID}}
		<p>Details:</p>
		<ul>
			{{if .Name}}<li>Name:{{.Name}}</li>{{end}}
			{{if .Surname}}<li>Surname:{{.Surname}}</li>{{end}}
			{{if .Age}}<li>Age:{{.Age}}</li>{{end}}
		</ul>
	{{else}}
		<p>Data not available</p>
	{{end}}
</html>
`

// 要用來替模板提供資料的結構
type Customer struct {
	ID      int
	Name    string
	Surname string
	Age     int
}

func Hello(w http.ResponseWriter, r *http.Request) {
	v1 := r.URL.Query() // 取得查詢參數
	customer := Customer{}

	id, ok := v1["id"]
	if ok {
		customer.ID, _ = strconv.Atoi(id[0])
	}

	name, ok := v1["name"]
	if ok {
		customer.Name = name[0]
	}

	surname, ok := v1["surname"]
	if ok {
		customer.Surname = surname[0]
	}

	age, ok := v1["age"]
	if ok {
		customer.Age, _ = strconv.Atoi(age[0])
	}

	// 建立名為 Example15.05 的模板，並填入 templateStr 模板字串用於解析
	tmp1, _ := template.New("Example15.05").Parse(templateStr)

	// 使用 customer 的資料填入模板，並將結果寫入 ResponseWriter （傳給客戶端）
	tmp1.Execute(w, customer)
}
func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
