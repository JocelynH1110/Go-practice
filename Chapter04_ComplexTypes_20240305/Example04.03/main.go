// 4-3-2  陣列的比較
// 如果有兩個陣列，都接受相同型別的元素，但元素的數量（長度）不一樣，則這兩個陣列就不相容，不能拿來比較。
// 即長度、元素型別不一致，就無法拿來比較。

// 練習、比較陣列是否相同
package main

import "fmt"

func compArrays() (bool, bool, bool) {
	var arr1 [5]int
	arr2 := [5]int{0}
	arr3 := [...]int{0, 0, 0, 0, 0}
	arr4 := [9]int{0, 0, 0, 0, 9}

	return arr1 == arr2, arr1 == arr3, arr1 == arr4
}
func main() {
	comp1, comp2, comp3 := compArrays()
	fmt.Println("[5]int == [5]int{0}  :", comp1)
	fmt.Println("[5]int == [...]int{0,0,0,0,0} :", comp2)
	fmt.Println("[5]int == [9]int{0,0,0,0,0} :", comp3)
}

//若你想在程式中比較集合，使用陣列會快得多。
//切片、map 就不能這樣比較，只能手動用迴圈走訪兩個集合，再逐一比對其元素值。
