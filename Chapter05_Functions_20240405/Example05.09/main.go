// 5-5-3  用自訂函式型別作為傳回值
package main

import "fmt"

func main() {
	add := calculator("+") //接收 calculator() 接收傳回的函式
	subtract := calculator("-")

	fmt.Println(add(5, 6))
	fmt.Println(subtract(10, 5))

	fmt.Printf("add()  型別：%T\n", add)
	fmt.Printf("subtract() 型別：%T\n", subtract)
}

func calculator(operator string) func(int, int) int {
	//根據使用者的引數傳回對應的函式
	switch operator {
	case "+":
		return func(i, j int) int {
			return i + j
		}
	case "-":
		return func(i, j int) int {
			return i - j
		}
	}
	return nil
}
