# 9-4 撰寫單元測試（unit test）
撰寫簡單的單元測試，並用 go test 工具來測試函式與套件。

### 撰寫測試檔
Go 語言測試檔的名稱不重要，但結尾必須加上「_test」，例如 shape_test.go。
在這檔案中，必須宣告一個測試用函式：
```go
func Test<名稱>(t *testing.T)
```
 
函式名稱也不重要，但必須以「Test」開頭，此函式會接收一個型別為 testing.T 的指標變數 t （來自 testing 套件）。
