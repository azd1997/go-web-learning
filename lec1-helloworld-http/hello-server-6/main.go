package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})
	// 增加一个超时函数，睡两秒然后在页面打印
	mux.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2*time.Second)
		w.Write([]byte("timeout"))
	})

	// 新建server
	server := &http.Server{
		Addr: ":4000",
		Handler:mux,
		WriteTimeout:2*time.Second,		// 设置写超时
	}
	
	log.Println("Starting HTTP server......")
	log.Fatal(server.ListenAndServe())
}

type helloHandler struct {}

// 实现ServeHTTP方法就实现了http.Handler接口
func (_ *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

// 仅仅是将server从http.ListenAndServe这个封装中抽出来并没有意义。
// 这里我们对自定义的server设置一个 写超时。

// 我们访问 localhost:4000/timeout将不会有任何消息打印。因为2秒之后server认为超时，关闭了与客户端（浏览器）的连接。