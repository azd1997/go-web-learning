package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// 模板函数
// 前面已经在模板中使用 eq ne lt le gt ge 六个判断大小相等的函数，他们是内置于模板引擎中的函数
// 我们也可以自定义自己的模板函数，通过 Funcs 方法
// Funcs方法通过传入一个template.FuncMap(底层实现是map[string]interface{})来将函数注入

// 标准库模板引擎还有许多用途的内置函数

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 创建模板
		tmpl := template.New("test")

		// 添加自定义函数
		tmpl = tmpl.Funcs(template.FuncMap{
			"add": func(a, b int) int {
				return a+b
			},
		})

		// 解析模板内容
		tmpl, err := tmpl.Parse(`
result: {{add 1 2}}`)

		if err != nil {
			fmt.Fprintf(writer, "Parse: %s", err)
			return
		}

		// 调用模板对象的渲染方法
		err = tmpl.Execute(writer, nil)
		if err != nil {
			fmt.Fprintf(writer, "Execute: %s", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

