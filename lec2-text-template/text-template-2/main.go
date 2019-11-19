package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// 在模板中渲染变量


func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 创建模板对象并解析模板内容
		tmpl, err := template.New("test").Parse("The value is: {{.}}")	// {{.}}中间的.默认指向根对象，就是Execute()方法的第二个参数
		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}

		// 获取 URL 参数的值	"localhost：4000/?val=123" URL查询参数
		val := request.URL.Query().Get("val")

		// 调用模板对象的渲染方法
		err = tmpl.Execute(writer, val)		// 将变量val渲染到io.writer
		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}
	})

	log.Println("Starting HTTP Server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}