package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	log.Println("Starting HTTP server......")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

// 1. http.HandleFunc将某个函数与一个路由规则绑定。这里将打印“hello world"的匿名函数与路由'/'（根路径）绑定
// 2. http.HandleFunc的参数2必须满足函数签名 func(w http.ResponseWriter, r *http.Request) { // do your thing here}
// 3. http.ListenAndServe()函数的作用是启动服务器，监听发送到指定地址和端口号的HTTP请求，参数2后面再讲。返回值为运行中的error信息。
// 		如果地址是localhost的话，也可以省略，写成http.ListenAndServe(":4000", nil)

// 4. 这里http.HandleFunc参数2使用了匿名函数写法，但当真正业务开发建议独立成具名函数，便于整理。改进后的代码见hello-server-2
