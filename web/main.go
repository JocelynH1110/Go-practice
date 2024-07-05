package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func resolvePort() string {
	env := os.Getenv("PORT")
	if env == "" {
		return ":3000"
	}
	return fmt.Sprintf(":%s", env)
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1>HOHOHO</h1>")
	})
	port := resolvePort()
	fmt.Printf("Listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
