package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	n := flag.String("name", "", "Your first name")
	i := flag.Int("age", -1, "Your age")
	b := flag.Bool("married", false, "Are you married?")
	flag.Parse()

	if *n == "" { // 若名字旗標值為空字串，代表使用者沒有加上該旗標，或未給值
		fmt.Println("Name is required!")
		flag.PrintDefaults() // 印出所有旗標的預設值
		os.Exit(1)           // 結束程式
	}
	fmt.Println("Name：", *n)
	fmt.Println("Age ：", *i)
	fmt.Println("Married：", *b)
}
