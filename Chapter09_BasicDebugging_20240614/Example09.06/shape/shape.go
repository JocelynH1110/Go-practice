// 沿用第八章的 shape 套件，來測試它對於不同形狀傳回的名稱及面積是否正確。
package shape

type Shape interface {
	area() float64
	name() string
}

type Triangle struct {
	Base   float64
	Height float64
}
type Rectangle struct {
	Length float64
	Width  float64
}

type Square struct {
	Side float64
}

func GetName(shape Shape) string { //修改過的新函式，傳回 shape 介面的名稱
	return shape.name()
}

func GetArea(shape Shape) float64 { //修改過的新函式，傳回 shape 介面的面積
	return shape.area()
}

func (t Triangle) area() float64 {
	return (t.Base * t.Height)
}

func (t Triangle) name() string {
	return "三角形"
}

func (r Rectangle) area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) name() string {
	return "矩形"
}

func (s Square) area() float64 {
	return s.Side * s.Side
}

func (s Square) name() string {
	return "正方形"
}

// 以上所有結構的方法（函式）仍維持小寫英文字母開頭，這是因為不想匯出這些功能給使用者使用。
