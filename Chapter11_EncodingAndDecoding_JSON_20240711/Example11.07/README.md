## 11-3-3 略過欄位
* omitempty
當希望沒有被賦值（保持在零值）的欄位，不要被編碼到 JSON 資料中時，可以在 JSON 標籤加入一個屬性 omitempty（略過零值），好讓零值欄位被 Mashel() 忽略。  

例子、
```
type book struct {
	ISBN          string `json:"isbn"`
	Title         string `json:"title"`
	YearPublished int    `json:"yearpub,omitempty"`
	Author        string `json:"author"`
	CoAuthor      string `json:"coauthor,omitempty"`
}
```
顯示結果：
```
{"isbn":"9933HIST","title":"Herry Potter","author":"J.K"}
```
> omitempty 和前面的 JSON 鍵名稱以逗點隔開，且不可以有空格。
> 若有空格，則 omitempty 失效，零值一樣會顯示。且不會傳回非 nil 的 error 值，除非用 go vet 工具來檢查（lesson 17）。  


* 其它標籤效果
JSON 標籤：
```go
type book struct {
	ISBN          string `json:"-"`         // 短折線
	Title         string `json:"title"`
	YearPublished int    `json:"yearpub,omitempty"`
	Author        string `json:""`              // 沒有鍵名稱
	CoAuthor      string `json:",omitempty"`    // 沒有鍵名稱
}
```
賦值給結構的變數：
```
	b := book{}
	b.ISBN = "9933HIST"     // 由於已指定略過，這不會出現在 JSON 中。
	b.Title = "Herry Potter"
    b.YearPublished = 2000
	b.Author = "J.K"
    // 沒有對 CoAuthor 賦職，因此會因 omitempty 被略過。
```
結果顯示：
```
{"title":"Herry Potter","yearpub":2000,"Author":"J.K"}
```

###  JSON 標籤的變化：
|:---                       |:---   |
|json:"<鍵名稱>"            | 加入此欄位並指定鍵名稱 |
|json:""                    | 加入此欄位並沿用欄位名稱 |
|json:"<鍵名稱>,omitempty"  | 若欄位非零值，加入此欄位並指定鍵名稱，否則略過 |
|json:",omitempty"          | 若欄位非零值，加入此欄位並沿用欄位名稱，否則略過 |
|json:"-"                   | 略過欄位 |
