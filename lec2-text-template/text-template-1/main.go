package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// 初识文本模板引擎
// go内置 text/template


func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 创建模板对象并解析模板内容
		tmpl, err := template.New("test").Parse("Hello world!")
		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}

		// 调用模板对象的渲染方法
		err = tmpl.Execute(writer, nil)
		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}
	})

	log.Println("Starting HTTP Server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}