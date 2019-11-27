package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// 模板中的作用域
// with/if/range代码块会形成独立的作用域，可以操作全局变量，但with/if/range它们的作用域内定义的变量不能被外部使用

func main() {

	tmplText1 := `
{{$name1 := "alice"}}
name1: {{$name1}}
{{with true}}
	{{$name1 = "alice2"}}
	{{$name2 := "bob"}}
	name2: {{$name2}}
{{end}}
$name1 after with: {{$name1}}
`

	tmplText2 := `
{{$name1 := "alice"}}
name1: {{$name1}}
{{with true}}
	{{$name1 = "alice2"}}
	{{$name2 := "bob"}}
	name2: {{$name2}}
{{end}}
$name1 after with: {{$name1}}
$name2 after with: {{$name2}}
`	// 在text1末尾追加了一句 {{$name2}} ，会报错，这个报错是在解析模板时出错
	// Parse: template: test:10: undefined variable "$name2" 返回给客户端

	tmplText3 := `
{{$name1 := "alice"}}
name1: {{$name1}}
{{with true}}
	{{$name1 := "alice2"}}
	{{$name2 := "bob"}}
	name2: {{$name2}}
{{end}}
$name1 after with: {{$name1}}
`	// 在text1的with内部做了改动{{$name1 := "alice2"}}，这样内部的name1和外部的不是一个，所以最后打印出来结果是
	// $name1 after with: alice

	// 避免因为变量未使用而报错
	tmplText1, tmplText2, tmplText3 = tmplText1, tmplText2, tmplText3

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 创建模板并解析
		tmpl, err := template.New("test").Parse(tmplText2)
		// 这里name1是全局变量，而name2是在with代码块的作用域定义的

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

