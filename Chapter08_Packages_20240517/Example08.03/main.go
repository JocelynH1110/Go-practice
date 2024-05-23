// 練習、建立一個能計算形狀面積的套件
package main

import "Example08.03/shape"

func main() {
	t := shape.Triangle{Base: 15.5, Height: 20.1}
	r := shape.Rectangle{Length: 20, Width: 10}
	s := shape.Square{Side: 10}
	shape.PrintShapeDetails(t, r, s)
}

// 在其他地方建立自己的專案時，不管有沒有使用套件，請養成習慣替它建立一個 go.mod
