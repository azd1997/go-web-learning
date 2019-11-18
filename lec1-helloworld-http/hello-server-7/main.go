package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})

	// 新建server
	server := &http.Server{
		Addr: ":4000",
		Handler:mux,
	}

	// 创建系统信号接收器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit		// 等待中断信号
		if err := server.Close(); err != nil {		// server.Close()调用之后会告诉服务器停止接收新请求，并在处理完当前已接收的请求后关闭服务器
			log.Fatal("Close server:", err)
		}
	}()
	
	log.Println("Starting HTTP server......")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {	// 这个错误类型是主动调用server.Close()后产生的，属于正常关闭
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}

type helloHandler struct {}

// 实现ServeHTTP方法就实现了http.Handler接口
func (_ *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

// 优雅地停止服务
// 结合捕捉系统信号Signal、goroutine和管道channel实现服务器优雅停止

//(base) eiger@eiger-ThinkPad-X1-Carbon-3rd:~/gopath-default/src/github.com/azd1997/go-web-learning/lec1-helloworld-http/hello-server-7$ go run main.go
//2019/11/18 07:22:50 Starting HTTP server......
//^C2019/11/18 07:23:07 Server closed under request