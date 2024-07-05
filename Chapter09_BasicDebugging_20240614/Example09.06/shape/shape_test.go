package shape

import "testing"

func TestGetName(t *testing.T) {
	// 測試用結構
	triangle := Triangle{Base: 15.5, Height: 20.1}
	rectangle := Rectangle{Length: 20, Width: 10}
	square := Square{Side: 10}

	if name := GetName(triangle); name != "三角形" {
		t.Errorf("%T 形狀錯誤：%v", triangle, name) // 傳回值錯誤時回報測試錯誤
	}
}
