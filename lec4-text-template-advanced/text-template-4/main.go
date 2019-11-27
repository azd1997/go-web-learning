package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

// 模板复用

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 创建模板
		tmpl := template.New("test")

		// 添加自定义函数
		tmpl = tmpl.Funcs(template.FuncMap{
			"join": strings.Join,

		})

		// 解析模板内容
		tmpl, err := tmpl.Parse(`
{{define "list"}}
	{{join . ", "}}
{{end}}
Names: {{template "list" .names}}
`)
		//Names:
		//        Alice, Bob, Cindy, David

		// 调用join其实就是调用strings.Join
		// 通过 define <名称> 定义了一个局部模板 list，以根对象 . 作为参数调用 join 模板函数
		// 通过 template <名称> <参数> 调用list模板，将 .names 作为参数传入，传入的参数将成为局部模板 list 的根对象 . 。

		// 模板复用最核心的概念： 定义、使用和传参

		if err != nil {
			fmt.Fprintf(writer, "Parse: %s", err)
			return
		}

		// 调用模板对象的渲染方法
		err = tmpl.Execute(writer, map[string]interface{}{
			"names": []string{"Alice", "Bob", "Cindy", "David"},
		})
		if err != nil {
			fmt.Fprintf(writer, "Execute: %s", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

