// 3-3-4   大數值
//若你所需的數值超過（或低於）int64 與 uint64 的極限，可以向內建的 math/big 套件求助。
//只要透過它，原本可以對一般整數做的大多數動作，都一樣能套用在大數值上。

// 練習、使用大數值
package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	intA := math.MaxInt64 //int 整數
	intA = intA + 1

	bigA := big.NewInt(math.MaxInt64)
	bigA.Add(bigA, big.NewInt(1))

	fmt.Println("MaxInt64:", math.MaxInt64)
	fmt.Println("Int	:", intA)
	fmt.Println("Big Int :", bigA.String())
}

// 3-3-5 位元組（byte）
/*
Go 語言裡的 byte 其實就是 uint8 的別名，uint8 是以 8 個位元儲存的正整數。
每一個位元（bit）代表一個二進位值，亦即開或關（1 或 0）。
電腦運算從很早期就開始採用以 8 個位元一組的「位元組」編碼。
8 個位元總共有 256 種可能的開關組合。uint8 0~255 ，故256種。
*/
