// 5-5 以函式為型別的參數

// 5-5-1  自訂函式型別
/*
函式在 Go 語言中也是一種型別，就像是整數、字串、布林值型別那樣。
這表示我們可以將函式當成引數，傳遞給其他函式；函式也可以傳回函式，甚至可以拿函式賦值給變數（如前面的閉包函式）。

如果想把函式當引數，得指明該接收參數的型別。函式型別的定義會包括他的參數與傳回值型別。
任何函式的參數與傳回值只要完全符合該自訂型別，就可以被視為該自訂型別。
這代表你能傳入多個不同的函式做為引數，只要他們的特徵都符合參數定義即可。

＊＊自訂函式型別：
1.type message func()	//定義了一個名為 message 的新函式型別，特徵為 func()。不具備輸入參數、也不提供傳回值。

2.type calc func(int,int) string	//定義了一個名為 calc 的函式型別，接受兩個整數型別參數、並傳回一個字串型別的值。這型別和前面的 func() 會是兩個不同的型別。
*/

// 5-5-2  使用自訂函式型別的參數
//以下為定義一個函式，使之符合某個自訂函式型別，並把它當成引數傳給另一個函式：
/*
package main

import "fmt"

type calc func(int, int) string

func main() {
	calculator(add, 5, 6) //把其他函式當引數
}

func add(i, j int) string {
	result := i + j
	return fmt.Sprintf("%d +%d =%d", i, j, result)
}

func calculator(f calc, i, j int) {
	fmt.Println(f(i, j))
}
*/

// 以下為示範如何將幾個不同的函式傳給 calculator()：
// 在這版本中，直接將自訂函式型別寫在 calculator() 的特徵內，而 add()、subtract() 函式符合這個型別，故可當 calculator() 的引數。
/*
package main

import "fmt"

func main() {
	calculator(add, 5, 6)
	calculator(subtract, 10, 5)
}

func add(i, j int) int {
	return i + j
}

func subtract(i, j int) int {
	return i - j
}

func calculator(f func(int, int) int, i, j int) {
	fmt.Println(f(i, j))
}
*/

// 練習、建立各種函式來計算薪資
package main

import "fmt"

type salaryFunc func(int, int) int

func main() {
	devSalary := salary(50, 2080, developerSalary)
	bossSalary := salary(150000, 25000, managerSalary)

	fmt.Printf("經理薪資：%d\n", bossSalary)
	fmt.Printf("程式設計師薪資：%d\n", devSalary)
}

// 薪資計算
func salary(x, y int, f salaryFunc) int {
	pay := f(x, y)
	return pay
}

func managerSalary(baseSalary, bonus int) int {
	return baseSalary + bonus
}

func developerSalary(hourlyRate, hoursWorked int) int {
	return hourlyRate * hoursWorked
}

//若將來需要額外計算新的薪資，只需要在建立一個新函式，使之符合 salary() 函式要求的輸入函式型別，就可以當引數傳入了。這樣的彈性使我們不必再更動 salary() 函式本身的寫法，就能創造出更多的功能。
