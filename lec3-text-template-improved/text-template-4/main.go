package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// 模板中使用迭代操作（range语句）
// template支持迭代操作，以方便直接在模板中对集合类型的数据进行处理和渲染
// 三种类型： array slice map 可以迭代

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 创建模板对象并解析模板内容
		tmpl, err := template.New("test").Parse(`
仅获取值：
{{range $name := .Names}}
	{{$name}}
{{end}}

同时获取索引和值：
{{range $i, $name := .Names}}
	{{$i}} {{$name}}
{{end}}

迭代map也是一样：
{{range $key, $value := .}}
	{{$key}}: {{$value}}
{{end}}
`)

		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}

		// 调用模板对象的渲染方法
		err = tmpl.Execute(writer, map[string]interface{}{
			"Names": []string{"Alice", "Bob", "Carol", "David"},
		})
		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}
	})

	log.Println("Starting HTTP Server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

// 函数的调用一般是带括号的，只不过 text/template 可以在语法上省略这个括号