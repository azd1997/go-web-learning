package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)

	log.Println("Starting HTTP server......")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

// 1. http.HandleFunc其实也是标准库提供的简便写法，其内部会将传入的参数2（hello函数）转化为类型 http.HanldeFunc，
// 而这个类型实现了http.Handler接口

// 2. 这也意味着我们可以自己实现一个结构体，让它实现 http.Handler接口，然后传到 http.Handle()函数，见hello-server-3
