// 例子、撰寫兩個測試函式：
package shape // 由於測試對象是 shape ，故宣告為同一套件

import "testing"

// 測試 shape 套件的 GetName() 函式
func TestGetName(t *testing.T) {
	// 測試用結構
	triangle := Triangle{Base: 15.5, Height: 20.1}
	rectangle := Rectangle{Length: 20, Width: 10}
	square := Square{Side: 10}

	if name := GetName(triangle); name != "三角形" {
		t.Errorf("%T 形狀錯誤：%v", triangle, name) // 傳回值錯誤時回報測試錯誤
	}
	if name := GetName(rectangle); name != "長方形" {
		t.Errorf("%T 形狀錯誤：%v", rectangle, name)
	}
	if name := GetName(square); name != "正方形" {
		t.Errorf("%T 形狀錯誤：%v", square, name)
	}
}

// 測試 shape 套件的 GetArea() 函式
func TestGetArea(t *testing.T) {
	// 測試用結構
	triangle := Triangle{Base: 15.5, Height: 20.1}
	rectangle := Rectangle{Length: 20, Width: 10}
	square := Square{Side: 10}

	if value := GetArea(triangle); value != 155.775 {
		t.Errorf("%T 面積錯誤：%v", triangle, value) // 傳回值錯誤時回報測試錯誤
	}
	if value := GetArea(rectangle); value != 200 {
		t.Errorf("%T 面積錯誤：%v", rectangle, value)
	}
	if value := GetArea(square); value != 100 {
		t.Errorf("%T 面積錯誤：%v", square, value)
	}
}

// 上面將測試用的結構寫在個別測試函式中，而非宣告在 shape 的套件層級變數，以免影響到 shape 套件本身。
// 測試函式會使用這些結構來測試 shape 公開函式的傳回值，看看結果是否跟已知的正確結果相符。
