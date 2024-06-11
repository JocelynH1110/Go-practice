package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1>HOHOHO</h1>")
	})
	fmt.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
