//1-4-2 算符簡寫法
/*
1.+=<值>：就地加上值
2.-=<值>：就地減去值
3.++：遞增1
4.--：遞減1
*/
package main

import "fmt"

func main() {
	count := 5

	count += 5
	fmt.Println(count)

	count++
	fmt.Println(count)

	count--
	fmt.Println(count)

	count -= 5
	fmt.Println(count)
	name := "Jocelyn"
	name += " Huang"
	fmt.Println("Hola,", name)
}
