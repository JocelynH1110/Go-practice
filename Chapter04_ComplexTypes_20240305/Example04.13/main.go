// 4-4-5  切片的隱藏陣列換置
/*
鑑於切片的可變性和其運作方式，切片是無法拿來相互比較的。若是比較則 Go 語言會丟錯誤給你。唯一能做的是拿切片和 nil 做比較。
切片是 Go 語言的一種特殊資料結構，本身不直接儲存值，而是在背景透過隱藏陣列存放資料。
切片本身只指向該隱藏陣列的指標，代表切片從哪裡開始，然後就只紀錄了切片的長度和容量。
這三種屬性（pointor、length、capacity）使得切片變成該隱藏陣列的窗格（window）。

同一個隱藏陣列也能被多重切片共用（你從切片建立的所有新切片,都會指向同一個原始陣列），但彼此窗葛布一定一樣大，也就是某些切片的值會比其他切片多。

切片本身不儲存值，底下也可能會共用陣列，這表示當你更改其中一個切片的元素時，實際上是在修改隱藏陣列，連帶改變其他切片看到的窗格。若沒注意這會在開發時遇到出乎預期的細微 bug。

更大的問題是，當切片需要擴充到超過其隱藏陣列的大小（窗格比隱藏陣列更大）時，Go 會建立出更大的陣列，把舊陣列的內容搬過去、再把切片的指標指向這個新陣列。
這種「陣列換置」的動作，就可能會導致不同的切片失去連結性。
之後若在更改其中一方的元素，改變就不會反應在另一個切片內。

若想複製一份切片，又要確保該切片能指向一個新隱藏陣列，不跟原來切片有連結，有兩種方式：
1.用 append() 把來源切片附加到另一個無關的新切片。
2.用 Go 內建函式 copy() 把來源切片複製到目標切片。

＊＊ copy(<目標切片>, <來源切片>)

在使用 copy() 函式時，Go 語言不會改變目標切片的大小，所以要確定目標切片有足夠的空間容納要複製的所有元素
*/

// 練習、觀察切片的連結行為
package main

import "fmt"

func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s3 := s1[:]
	s1[3] = 99
	return s1[3], s2[3], s3[3]
}

func noLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copied := copy(s2, s1)
	s1[3] = 99
	return s1[3], s2[3], copied
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append([]int{}, s1...)
	s1[3] = 99
	return s1[3], s2[3]
}

func main() {
	l1, l2, l3 := linked()
	fmt.Println("有連結：", l1, l2, l3)

	n11, n12 := noLink()
	fmt.Println("無連結：", n11, n12)

	c11, c12 := capLinked()
	fmt.Println("有設容量，有連結：", c11, c12)

	cn11, cn12 := capNoLink()
	fmt.Println("有設容量，無連結：", cn11, cn12)

	copy1, copy2, copied := copyNoLink()
	fmt.Print("使用 copy() ，無連結：", copy1, copy2)
	fmt.Printf("(複製了 %v 個元素)\n", copied)

	a1, a2 := appendNoLink()
	fmt.Println("使用 append()，無連結：", a1, a2)
}
