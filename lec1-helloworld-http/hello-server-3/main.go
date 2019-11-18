package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", &helloHandler{})

	log.Println("Starting HTTP server......")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

type helloHandler struct {}

// 实现ServeHTTP方法就实现了http.Handler接口
func (_ *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
