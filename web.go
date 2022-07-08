package main
import (
	"fmt"
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/", index) // index 为向 url发送请求时，调用的函数
	log.Fatal(http.ListenAndServe("127.0.0.1:8888", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "fwt测试")
}