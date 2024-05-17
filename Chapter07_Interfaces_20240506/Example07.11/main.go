//練習、分析介面資料
/*
會有一個 map ，其索引鍵是字串，對應值則是 interface{} 空介面，也就是說它可能儲存不同型別的資料。任務是找出每一個鍵對應元素值的型別，並把每個元素包裝成結構 record 。record 中會有不同的欄位來紀錄原始的鍵與值，並有個字串欄位來紀錄該資料的型別。
*/
package main

import "fmt"

type person struct {
	lastName  string
	age       int
	isMarried bool
}

type animal struct {
	name     string
	category string
}

type record struct {
	key       string
	data      interface{}
	valueType string //值的型別
}

func main() {
	m := make(map[string]interface{}) //建立並初始化一個 map

	//在 map 內加入不同型別的多筆資料
	m["person"] = person{lastName: "Doe", isMarried: false, age: 20} //初始完畢後加入值，map名稱[索引鍵]=值
	m["firstname"] = "Smith"
	m["age"] = 50
	m["isMarried"] = true
	m["animal"] = animal{name: "Sala", category: "Cat"}

	//解析 map 每個元素，轉換成 record 結構和放進一個切片
	rs := []record{}
	for k, v := range m {
		rs = append(rs, newRecord(k, v))
	}

	//印出 record 結構切片的內容
	for _, v := range rs {
		fmt.Println("Key : ", v.key)
		fmt.Println("Data : ", v.data)
		fmt.Println("Type : ", v.valueType)
		fmt.Println()
	}
}

func newRecord(key string, i interface{}) record {
	r := record{}
	r.key = key
	switch v := i.(type) { // 對 map 元素值做型別 switch
	case int:
		r.valueType = "int"
		r.data = v
		return r
	case bool:
		r.valueType = "bool"
		r.data = v
		return r
	case string:
		r.valueType = "string"
		r.data = v
		return r
	case person:
		r.valueType = "person"
		r.data = v
		return r
	default:
		r.valueType = "unknown"
		r.data = v
		return r
	}
}

/*
若單純要用字串紀錄某個值的型別，可以這樣寫：
r.valueType = fmt.Sprintf("%T",i)

fmt.Sprintf() 的作用和 fmt.Printf() 很像，但不會印出字串到主控台，而是傳回格式化後的字串。
*/
