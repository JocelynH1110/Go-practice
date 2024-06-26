// 6-3 其他程式語言的錯誤處理方式
/*
Go 語言處理錯誤的方式和 Java、Python、C#、Ruby 等語言都不一樣，這些其他語言做的是所謂的「例外處理」（exception handling）

在大部分語言程式中，例外處理是隱性（implicit）的：
任何函式都有可能出錯和拋出例外，但事先無法預知誰會這麼做，只能試著用 try...error 來攔截他們。若沒有處理例外，該函式會導致整個程式當掉。

Go 語言，錯誤的處理是顯性（explicit）的：
許多函式會很明確傳回一個你無法拒絕的錯誤值，但該值在函式執行成功時是 nil 。
就算真的傳回非 nil 錯誤，函式也不見得會讓程式當掉，但有責任要在錯誤值是 nil 時處裡它。

大部分程式語言中，若某些功能有發生錯誤的可能性，就得撰寫 try...catch 敘述來包住它。
但在 Go 語言裡，你會很明確的先接收 error 值，然後在自己判斷要對它做什麼事：

var,err := someFunc()
if err != nil{	//若有錯誤存在，做些處理然後再把它繼續往外傳
	return err
}
return nil	//沒有錯誤，對上一層傳回 nil

若只是要檢查一個函式是否正確執行，也可以寫成如下：（更簡潔一點）
if _,err:=somrFunc();err != nil{

}
*/
