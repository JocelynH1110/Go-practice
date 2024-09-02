package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 從請求主體取出名稱為 myFile 的檔案（multipart.File 型別）
	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// multipart.File 符合 io.Reader 介面，故可用 io.ReadAll() 讀取內容
	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// 將檔案寫入伺服器端的系統
	err = os.WriteFile(fmt.Sprintf("./%s", fileHeader.Filename), fileContent, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// 顯示並回傳已上傳檔案的訊息
	log.Printf("%s uploaded", fileHeader.Filename)
	w.Write([]byte(fmt.Sprintf("%s uploaded!", fileHeader.Filename)))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
