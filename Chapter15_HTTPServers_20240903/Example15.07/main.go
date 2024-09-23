package main

import (
	"log"
	"net/http"
)

func main() {
	// 對任何路徑提供 index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	// 將 /statics 路徑對應到本地的 /public 資料夾
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("./public"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
