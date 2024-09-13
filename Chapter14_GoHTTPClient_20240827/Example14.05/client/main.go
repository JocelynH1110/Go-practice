package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func postFileAndReturnResponse(filename string) string {
	fileDataBuffer := bytes.Buffer{}                 // 建立一個 buffer
	mpWriter := multipart.NewWriter(&fileDataBuffer) // 建立 multipart.Writer

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 用 multipart.Writer 建立準備傳送的 MIME 檔案
	formFile, err := mpWriter.CreateFormFile("myFile", file.Name())
	if err != nil {
		log.Fatal(err)
	}

	// 將原始檔案的內容拷貝到 MIME 檔案
	if _, err := io.Copy(formFile, file); err != nil {
		log.Fatal(err)
	}
	mpWriter.Close() // 關閉 multipart.Writer （必要）

	// 用 POST 請求送出 MIME 檔案並讀取回應
	// 使用 multipart.Writer 來指定標頭內的內容類型為 multipart/form-data
	r, err := http.Post("http://localhost:8080", mpWriter.FormDataContentType(), &fileDataBuffer)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	data := postFileAndReturnResponse("./test.txt")
	fmt.Println(data)
}
