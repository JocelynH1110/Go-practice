// 1-2-4 當型態錯誤時
package main

import "math/rand"

func main() {
	var seed int64 = 123456789 //rand.seed 是 int64 的型態，如果這裡少打型態就會產生錯誤，因為這一串數字是 int。
	rand.Seed(seed)
}
