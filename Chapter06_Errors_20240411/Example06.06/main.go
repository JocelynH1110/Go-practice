// 6-4-4  使用 fmt.Errorf() 建立 error 值
/*
前面是用 errors.New() 建立 error 值。
另一個方式是使用 fmt.Errorf() ，讓你建立格式化的錯誤訊息：
*/

func payDay(hoursWorked,hourlyRate int)(int,error){
	if hourlyRate < 0 || hourlyRate > 75{
		return 0, fmt.Errorf("無效的時薪： %d",hourlyRate)
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		return 0, fmt.Errorf("無效的一週工時： %d",hoursWorked)
	}
}

// 這表示我們更可以把其他 error 值的內容讀出來，連同其他訊息合併成一個新的 error 值：
func payDay(hoursWorked,hourlyRate int)(int,error){
	if hourlyRate < 0 || hourlyRate > 75{
		return 0, fmt.Errorf("payDay 錯誤： %s",ErrHourlyRate.Error())
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		return 0, fmt.Errorf("payDay 錯誤： %s",ErrHourlyWorked.Error())
	}
}
這種合併法意味著舊的 error 值會被新型別蓋掉，而特定的 error 值可能是有其特殊意義的。
此外，原本的 error 值也可能擁有額外的欄位、方法等等，而這些資訊都會在合併過程中消失。

因此在 Go 1.13 起，fmt.Errorf() 提供了另一種結合 error 值的做法 — error 值可以包覆（wrap）其他的 error 值。
辦法是在格式化字串中用 %w 符號來對應到要被包覆的 error：
func payDay(hoursWorked,hourlyRate int)(int,error){
	if hourlyRate < 0 || hourlyRate > 75{
		return 0, fmt.Errorf("無效的時薪： %w",ErrHourlyRate)
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		return 0, fmt.Errorf("無效的一週工時： %w",ErrHoursWorked)
	}
}

觀察 error 套件的原始碼，會發現 fmt.Errorf() 傳回的 error 型別實際上為 wrapError 結構，它有個 error 型別欄位能記住它「包覆」的錯誤值，並有額外的方法能夠讀取該欄位：
type wrapError struct{
	msg string
	err error
}

//wrapError 有實作 Error() 因此符合 error 因此符合介面型別
func (e *wrapError) Error() string{
	return e.msg
}

func (e *wrapError) Unwrap() error{
	return e.err
}

當你在 fmt.Errorf() 使用 %w 符號來包覆其他 error 值時，後者會被儲存到新 error 值（wrapError 型別）的 err 欄位中：
這也意味著當函式接受了 error 值並往外傳時，可以將錯誤值一層層包覆起來形成所謂的錯誤鏈（error chain），而當中所有資訊都不會喪失。

Go 語言也提供了兩個新方法，error.Is() 和 error.As() ，使你能檢查錯誤鏈中是否存在著某個 error 值或特定的 error 別：
⊕ error.Is(err,target) 能檢查一個 err （一個 error 值）中是否包含 target 的值，有的話傳回 true 。
⊕ error.As(err,target) 能檢查一個 err （一個 error 值）中是否有 error 值的型別符合 target 的型別，有的話就該將值指定給 target（target 需為指標）並傳為 true 。
簡言之，errorIs() 是在做值的檢查，error.As() 的效果則像是 error 型別斷言。
