## 8-3-5 下載第三方模組或套件
若要使用其他人已經放置在網路上的公開套件，可用 go get 指令下載到系統的 $GOPATH\pkg\mod 位置。
必要時，可用 go mod tidy 來重整 go.mod 檔案，讓 Go 語言尋找你下載的套件位於何處。

### 用 go get 下載第三方模組或套件
我們使用 Go 語言官方提供的範例模組 example （http://github.com/golang/example），當中的 stringutil 套件，內含一個可反轉字串的 Reverse() 函式。

打開連結會看到下面的安裝說明：
```
go get golang.org/x/example/hello
```

go get 後的網址，代表下載對象是 golang.org/x/example 這個模組以下的 hello 套件。

* 主要流程：
1. 在主程式 import 第三方套件
2. 在專案路徑下執行（套件名稱可自取）
`go mod init 套件名稱`
3. 下載套件：
`go get golang.org/x/example/stringutil`
OR
`go get github.com/ozgio/strutil`
4. 重整 go.mod：
`go mod tidy`

### 用 go mod tidy 整理/更新 go.mod
如果 go.mod 沒有自動更新，可至主控台執行 go mod tidy 來重整。
專案多出來一個 go.sum 檔案，是用來紀錄套件的雜湊長度（hash），以便確保下載的套件未經竄改。


補充、
`go mod vendor`
在主控台執行這個指令，這會將第三方模組/套件拷貝一份放在專案資料夾的 vendor 子目錄下，系統內沒有的話則會嘗試自行下載。
若專案內含有 vendor 資料夾，專案就會使用它內含的原始碼，而不是從 $GOPATH 存取。
這樣做的好處之一是能將相依套件打包在一起，其他人不需下載就能直接使用，缺點是套件或模組本身可能很占空間。

