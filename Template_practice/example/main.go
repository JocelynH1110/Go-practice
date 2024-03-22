package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var md = `Hello,{{ . }}`

func main() {
	tpl := template.Must(template.New("first").Parse(md)) //創建一個名為 first 的 template，並用此 template 進行parse 進行解析模板。
	if err := tpl.Execute(os.Stdout, "Jocelyn"); err != nil {
		log.Fatal(err)
	}
	fmt.Print("\n")
}
