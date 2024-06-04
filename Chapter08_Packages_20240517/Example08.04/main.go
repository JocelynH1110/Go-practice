package main

import (
	"fmt"

	//以下為兩種方式匯入第三方套件
	"github.com/ozgio/strutil" //第三方套件路徑
	//"golang.org/x/example/stringutil"  匯入第三方套件 （但這個找不到 stringutil 這個套件）
)

func main() {
	// 呼叫套件功能來反轉字串
	fmt.Println(strutil.Reverse("!selpmaxe oG,ollaH"))
}

//在 go run . 之前都要先在主控台 go get 上面兩種 import 提供的其中一種，然後 go mod tidy
//但也可以先在 import 裡打上路徑套件，再到主控台用 go mod tidy 重整 go.mod. 這樣 Go 語言會自動尋找並下載你指定的套件以及其相依的套件。
