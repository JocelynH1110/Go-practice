//Chapter 1 變數與算符

package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{
	"你好",
	"Ciao",
	"こんにちは",
	"안녕하세요",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(helloList))
	msg, err := hello(index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil
}
