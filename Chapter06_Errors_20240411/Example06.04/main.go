// 6-4-2  error 型別定義
/*
繼續研究 Go 語言標準套件中的 error 型別。從函式庫中的 errors.go 著手： 
type errorString struct{
	s string
}

這結構型別位於 error 套件中，有個字串型別欄位 s 來儲存錯誤的內容。
errorString 型別和 s 都是以英文小寫開頭，代表他們是不可匯出或公開的，亦不能在外部程式直接使用他們。（lesson 8）

func (e *errorString) Error() string{
	return e.s
}
