go get :
 - github.com/tidwall/gjson
 - github.com/gin-gonic/gin

run :
 - go run server.go routes.go

tests :
 - go test ./... -cover -bench=.

goos: linux
goarch: amd64
pkg: ml/api
BenchmarkPrice-8   	       3	 459731130 ns/op
PASS
coverage: 80.0% of statements
ok  	ml/api	2.902s
goos: linux
goarch: amd64
pkg: ml/itemCat
BenchmarkLoad-8       	       2	 554593959 ns/op
BenchmarkGetPrice-8   	    5000	    284429 ns/op
PASS
coverage: 86.4% of statements
ok  	ml/itemCat	8.151s
goos: linux
goarch: amd64
pkg: ml/math
BenchmarkStats-8   	50000000	        26.5 ns/op
PASS
coverage: 93.8% of statements
ok  	ml/math	1.354s
