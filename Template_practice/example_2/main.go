package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var md = `個人信息：
	姓名：{{ .Name}}
	年齡：{{ .Age}}
	愛好：{{ .Hobby -}}  
`

//-}}表示右邊的空格應被去除
//{{- 10}} 表示向左刪除空格，印10
//{{-10}}  表示印10

type People struct {
	Name string
	Age  int
}

func (p People) Hobby() string {
	return "唱、跳、rap"
}

func main() {
	tpl := template.Must(template.New("first").Parse(md))
	p := People{
		Name: "Kiki",
		Age:  30,
	}
	if err := tpl.Execute(os.Stdout, p); err != nil {
		log.Fatal(err)
	}
	fmt.Print("\n")
}
