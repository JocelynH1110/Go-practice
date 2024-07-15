package main

import (
	"encoding/json"
	"fmt"
)

type book struct {
	ISBN          string `json:"isbn"`
	Title         string `json:"title"`
	YearPublished int    `json:"yearpub"`
	Author        string `json:"author"`
	CoAuthor      string `json:"coauthor"`
}

func main() {
	b := book{}
	b.ISBN = "9933HIST"
	b.Title = "Herry Potter"
	b.Author = "J.K"
	// 沒有對 YearPublished 和 CoAuthor

	json, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", json)
	fmt.Println(string(json))
}
