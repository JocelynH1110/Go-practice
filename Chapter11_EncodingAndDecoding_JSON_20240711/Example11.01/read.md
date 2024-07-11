# 11-1 前言
Go 語言要處理的資料不見得都是來自程式本身，也有可能會和外界交換資料。  而最常見的資料交換格式之一，就是 JSON 。

## JSON 格式資料
* JSON（JavaScript Object Notation, JavaScript 物件表示法）
* 儘管來源於 JavaScript，如今已被許多程式語言用來儲存和交換資料。
* 也常用於 HTTP 伺服器和客戶端之間的通訊（lesson 14、15）
* 也有靜態網站會拿 JSON 來產生網頁。
* NoSQL 伺服器等技術也採用 JSON 作為儲存格式。

JSON 是一種與任何程式語言無關的純文字格式，設計宗旨為精簡至上，不若 XML 繁瑣。且自帶描述資訊，提高其可讀性並降低了撰寫難度。  

* JSON 具備以下特性：
	* 輕量（lightweight）
    * 和程式語言無關（programming language-agnostic）
    * 自我描述（self-describing）
    * 使用鍵與值對（key/value pairs）

諸如 RESTfil API 這類網路服務，之所以會採用 JSON 而非 XML 做為資料交換格式，就是因為其簡明輕巧，且更容易閱讀。

> ### JSON 鍵與值寫法
```
// JSON
{
    "firstname":"Jocelyn"
    "lastname":"Huang"
}

// XML
<avenger>
    <firstname>Jocelyn</firstname>
    <lastname>Huang</lastname>
</avenger>

```
> 鍵：用雙引號刮起來的字串。  
> 值：可以次是多種資料型別。  
> 鍵與值之間以：連接。  
> 鍵與值不只有一組，各組織間以逗號隔開。

> ### JSON 的值也可以是陣列：（用中括號）
```
{
    "phonenumbers":["123-123-123","111-111-111"]
}
```

### JSON 的值也可以是另一筆 JSON 資料（或 JSON 物件）：
```
{
    "phonenumbers":[
        {"type":"business","number":"123-123-123"},     // JSON 物件
        {"type":"home","number":"111-111-111"},     // JSON 物件
    ]
}
```

### JSON 可用的值型別：

|  型別   | 範例  |
|  :----  | :----  |
| string  |  {"firstname":"Nana"} |
| number  |  {"age":30} }         |
| boolean |  {"ismarried":false}  |
| array   |  {"hobbies":["Go","Saving Earth","Shield"]} |
| null    |  {"middlename":null}  |
| object  |  另一筆 JSON 資料     |

