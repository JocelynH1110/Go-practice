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
