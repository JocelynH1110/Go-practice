package main

import (
	"errors"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "log:", log.Ldate|log.Lmicroseconds|log.Llongfile)
	logger.Println("Start of our app")
	err := errors.New("application aborted!")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("End of our app")
}
