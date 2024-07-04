// 用追蹤訊息找出錯誤發生位置
// 可以在 main() 印出隨機數 r 的值，以判斷是哪個函式傳回 error，但在兩個函式內部印出訊息會更顯著簡單判別。

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	r := random(1, 20)
	err := a(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = b(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func random(min, max int) int {
	// rand.Seed(time.Now().UTC().UnixNano()) 產生為隨機資料，Go 1.20 以後會自動產生
	return rand.Intn((max-min)+1) + min
}

func a(i int) error {
	if i < 10 {
		fmt.Println("錯誤發生在a()")
		return errors.New("incorrect value")
	}
	return nil
}

func b(i int) error {
	if i >= 10 {
		fmt.Println("錯誤發生在b()")
		return errors.New("incorrect value")
	}
	return nil
}
