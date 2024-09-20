# 15-4 使用模板產生網頁
由前一章得知，網頁伺服器除了能傳回文字資料，也能分享 JSON 這種結構化資料。  
然而 JSON 主要用途是讓程式交換資料，它得解析和處理過才能轉換成適合瀏覽的網頁。  
前一章最後一個練習題用了 fmt.Sprintf() 來產生格使化字串，但它仍難以應付更複雜的動態資料。  

如果網頁的內容格式是固定或有跡可循的，只有資料（例、使用者透過 URL 傳入名字等）是動態的，那還可以使用一個新技巧，叫 **網頁模板（template）**。  

模板：就是一串文字構成的骨架，當中有些部份留白，讓模板引擎抓一些值和填進去，如下：
```
Hello there my name is {{name}} 
                                    ---> 模板引擎 ---> Hello there my name is Bill 
{{name}} = Bill
```
> 上面的左上角是模板，{{name}} 是填空處。當把 "Bill" 值傳給模板引擎，填空處就會被換成其值，產生出動態內容。


在 Go 語言裡，提供了兩種模板套件：
1. 用於文字（text/template）
2. 用於 HTML （html/template）  

既然我們是在用 HTTP 伺服器產生網頁，下面就會用 HTML 模板套件，但它操作起來跟文字模板套件是一樣的。

* html/template 能防範跨網站指令攻擊
兩者模板套件的差異在於，html/template 會套用自動字元跳脫（autoescape），也就是將符合 HTML、CSS、JavaScript 等指令的特殊字元轉換過，以免被用於「跨網站指令碼」（cross-site scripting，XSS）攻擊。  

例如、一個網站讓使用者輸入姓名，然後直接填入模板的填空處。此時攻擊者可以故意填入 JavaScript 碼來使之執行：
```html
<script>alert("XSS attack!")</script>
```

若模板需要的資料是由 URL 參數提供，那攻擊者更可藉由提供釣魚網址的方式來夾帶程式碼，藉此竊取其他使用者瀏覽網站時填入的個資等。  
使用 html/template 套件能有效防堵這類攻擊，使值入的任何程式碼都不會被執行。  

Go 的標準模板套件已經很好用，將來仍可考慮用外部套件（如、Hero）來大幅提昇產生效能。  


### 使用 html/template 套件
1. Go 語言的 HTML 模板套件提供一種模板語言，讓我們能像這樣單純取代填空處的值：
`{{name}}`

2. 也能用模板語言進行複雜一點的條件判斷：
`{{if age}} Hello {{else}} Bye {{end}}`
> 若 age 內容不為 nil，模板引擎就會填入字串 "Hello"，反之則使用 "Bye"。條件判斷必須用{{end}}收尾。

3. 模板變數也不見得只能是簡單的數字或字串，更可以是物件。  
例、我們有個結構，內含一個欄位叫 ID，就能像下面一樣把該欄位填入模板：
`{{.ID}}`
> 這樣很方便，這表示我們能只傳一個結構給模板，而不是得傳一堆個別的變數。

練習、套用 HTML 模板
本練習目的是用模板來打造結構更好的網頁，而其內容是透過 URL 的 QueryString 傳入的。  
下面程式中，會顯示消費者的一些基本資訊：
1. ID
2. Name
3. Surname
4. Age

故查詢該網頁時，完整的 URL 會如下：
`http://localhost:8080/?id=代碼&name=姓名&surname=姓氏&age=年齡`

為了簡化起見，就算使用者輸入多重參數，程式也只會讀取第一項。  
若未提供 id ，那頁面只會顯示「資料不存在」；其他三項資料，缺少的項目會直接隱藏。

```go
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
```
> Execute() 方法，第一個參數接收 io.Writer 介面型別，而 http.ResponseWriter 就符合。於是填入值的模板字串（一個 HTML 網頁）就會被傳給客戶，並在瀏覽器中顯示出來。
> 讀取 ID 和 Age 時，呼叫了 strconv.Atoi() 來將字串轉成數字。若轉換錯誤，第二個參數會傳回 error。理論上要處理錯誤，但這裡忽略了，因為輸入錯誤會就會得到零值，故在這不希望因此讓伺服器掛掉。


這程式效率有點不太好，因為每次處理請求時都會呼叫 template.New() 來產生一個新模板。更適當的作法是把模板存在一個請求處理器結構中，然後在初始化時產生一次就好。
