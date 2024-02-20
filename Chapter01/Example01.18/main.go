//1-7	常數（constants）

//const <常數> <型別> = <值>

/*
常數就像是變數，但無法改變他的初始值。
如果程式執行時，有個數字不需變動、也不該改變時，常數就能派上用場。

常數宣告跟使用 var 類似，但改用 const 關鍵字。
宣告常數時，初始值必不可少，型別可有可無，若沒指定型別，GO 語言會自己推斷。
初始值可以是值或一段簡單的運算式，甚至可以引用其他常數。
*/
package main

import "fmt"

const GlobalLimit = 100                  //單筆資料上限
const MaxCacheSize int = 2 * GlobalLimit //快取最大容量

const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string //快取集合

func cacheGet(key string) string { //從快取取出某個鍵的值
	return cache[key]
}

func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val //寫入資料
}

func SetBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

func SetCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}

func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

func main() {
	cache = make(map[string]string) //初始化快取

	//在快取寫入資料
	SetBook("1234-5678", "Get Ready To Go")
	SetCD("1234-5678", "Get Ready To Go Audio Book")

	//讀取和印出快取資料
	fmt.Println("Book：", GetBook("1234-5678"))
	fmt.Println("CD：", GetCD("1234-5678"))
}

//當寫入資料達到常數值設定的200筆時，快取表就不會再接受任何資料。若要改變快取表中的前綴詞或快取大小，直接改常數定義即可
