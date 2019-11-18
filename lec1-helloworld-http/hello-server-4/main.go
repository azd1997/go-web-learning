package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})

	log.Println("Starting HTTP server......")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

type helloHandler struct {}

// 实现ServeHTTP方法就实现了http.Handler接口
func (_ *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

// 这一个例程，使用ServeMux。
// 前面直接使用 http.Handle或者HandleFunc其实是对一个默认的http.ServeMux对象 DefaultServeMux 的Handle方法做了封装
// ServeMux的作用就是服务复用器Multiplexer
// 这节新建一个ServeMux来实现Hello

// 本质上就是一个带有路由层的http.Handler具体实现，并以此为基础提供大量辅助方法