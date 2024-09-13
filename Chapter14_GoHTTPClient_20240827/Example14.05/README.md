## 14-4-2 用 POST 請求上傳檔案
* 常見範例：當想從本地電腦用 POST 請求上傳一個檔案到伺服器上。如、上傳照片等。  
* 為此，我們得在 POST 請求標頭中指定使用 **MIME（多用途網際網路郵件擴展）** 格式來傳送檔案。  
* 這種標準能將檔案切割成較小的訊息以利傳送，且支援多媒體類型。

上傳檔案比上傳單純表格資料更複雜：  
客戶端必須將要上傳的檔案轉換成 MIME 格式，伺服器收到後也要讀取它，再將其寫入到系統中。
Go 語言標準套件 mime/multipart 可以替我們應付 MIME 物件的建立，並傳回適當的請求標頭（multipart/form-data）。  


練習、用 POST 請求將檔案傳給伺服器

* 伺服器程式：下面是簡單的伺服器，會從主體讀取使用者傳送的 MIME 檔案，並將其內容寫入到系統中。
```go
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
```
> 使用了 os.WriteFile() 函式，它能建立一個新檔案，並將一個 []byte 切片寫入到其內容，不需另外開啟檔案物件。

* 客戶端程式：如何建立 MIME 檔案來給伺服器解讀。
```go
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
```
* 執行程式
1. 先執行 server：
```shelll
Chapter14_GoHTTPClient_20240827/Example14.05/server$ go run .
```

2. 接著在 client 子目錄下新增一個 test.txt，在執行 client 裡的 main.go：
```shell
Chapter14_GoHTTPClient_20240827/Example14.05/client$ go run .
test.txt uploaded!
```

3. 回到 server 子目錄，資料夾下會多一個 test.txt ，顯示檔案上傳成功，伺服器的主控台也會顯示上傳訊息：
```shell
Chapter14_GoHTTPClient_20240827/Example14.05/server$ go run .
2024/09/02 15:39:11 test.txt uploaded
```
