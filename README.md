# kevctl

# Installing
Using Cobra is easy. First, use `go get` to install the latest version
of the library.     

```
git clone git@github.com:kaidlog/kevctl.git
```

Next, include Cobra in your application:
Compile `main.go`
```go
go build main.go
```

# Feature

1. Json 轉換和處理
```
go run main.go json
```
2. sql轉換和處理
```
go run main.go json
```
3. 時間格式處理
```
// 計算當下時間
go run main.go time now
```
```
// 計算現在時間 + 兩小時
go run main.go time calc -d=2h
```
```
// 指定時間 + 兩小時
go run main.go time calc -c="2029-09-0412:02:33" -d=2h
```
```
4. 單詞格式轉換
```
go run main.go word -m=1 -s=jack
```
