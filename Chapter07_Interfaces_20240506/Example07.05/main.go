// 練習、使用多型來計算不同形狀的面積
// 能印出圓形、三角形、正方形三角形、正方形的名稱與面積。負責印出資訊的函式會接收 Shape 這個介面型別的數量不定參數，使任何滿足 Shape 規範的形狀都可以當成引數傳入。
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Name() string
	Area() float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	s := square{side: 10}
	c := circle{radius: 3.5}
	t := triangle{base: 3.6, height: 4.1}
	printShapeDetails(s, c, t)
}

func printShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("%s 的面積：%.2f\n", item.Name(), item.Area())
	}
}

// 以下是實作介面所需的方法
func (c circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) Name() string {
	return "圓形"
}

func (s square) Area() float64 {
	return s.side * s.side
}

func (s square) Name() string {
	return "正方形"
}

func (t triangle) Area() float64 {
	return (t.base * t.height) / 2
}
func (t triangle) Name() string {
	return "三角形"
}

//每一種形狀都滿足 Shape 介面，因為他們都具有 Area()、Name() 兩種方法，且方法特徵也吻合。儘管每個結構的欄位有所不同，他們都可以被 printShapeDetails() 函式使用。
