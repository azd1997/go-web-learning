package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})

	// 新建server
	server := &http.Server{
		Addr: ":4000",
		Handler:mux,
	}
	
	log.Println("Starting HTTP server......")
	log.Fatal(server.ListenAndServe())
}

type helloHandler struct {}

// 实现ServeHTTP方法就实现了http.Handler接口
func (_ *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

// 这一个例程，使用自定义的http.Server。
// http.ListenAndServe()函数其实内部就先创建了一个http.Server，然后调用server.ListenAndServe()
// 我们也可以自己自定义一个Server，并且自定义的程度非常高，包含了go标准库提供的所有可能的选项，包括坚挺地址、服务复用器和读写超时。。。
