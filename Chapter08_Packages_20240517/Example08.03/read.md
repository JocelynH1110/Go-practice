# 8-3 管理套件

Go 編譯器會去哪裡找我們應用程式所引用的套件？

## 8-3-1 GOROOT
Go 編譯器必須知道如何找到我們用 import 匯入的套件的原始檔，才能建置和安裝套件。
$GOROOT 就是 Go 語言在電腦裡的安裝路徑。

＊查詢所有 Go 環境變數
```shell
$ go env
```

＊查詢 GOROOOT 變數
```
$ echo $GOROOT
```

## 8-3-2 GOPATH
$GOPATH 通常指向使用者家目錄下的 Go 目錄。

在 $GOPATH 下通常有三個子目錄：
1.bin：當執行 go install 命令時，Go 語言會把編譯好的二進位執行檔放在此處。
2.pkg：用來放編譯過的套件，也會於其 mod 子資料夾下存放用 go get 下載的第三方套件。
3.src：在 GO 1.11 版前，使用者所有專案和套件都得置於 $GOPATH\src 目錄下，但新的 Go Module 功能解除了這種限制。

## 8-3-3 Go Module
Go 1.11 版起，新的 Go Modules 功能取代了 $GOPATH ，這功能從 Go 1.16 版後也預設為啟用。
模組（module）代表一系列套件的集合，而模組路徑（module path）會被用來協助 Go 語言尋找你的套件，如此一來就不用仰賴 $GOPATH 來放置套件了。

若有使用到自訂或外部套件，必須在專案的根目錄建立一個 go.mod 檔案：

＊建立 go.mod 指令
```shell
go mod init 模組名稱
```

模組名稱不需跟專案名稱或專案資料夾同名。

當對程式加入或移除了套件時，應該在專案目錄下輸入以下指令，以重整 go.mod 的內容：
```
go mod tidy
```

Go Module 不只能用來尋找套件，也能用於套件版本控管。
