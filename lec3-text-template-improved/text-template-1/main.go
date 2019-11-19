package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// 模板中定义变量


func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 创建模板对象并解析模板内容
		tmpl, err := template.New("test").Parse(`
{{$name := "Alice"}}
{{$age := 18}}
{{$round2 := true}}{{$round2 := false}}
Name: {{$name}}
Age: {{$age}}
Round2: {{$round2}}
`)		// 这里在模板中定义了三个变量并使用。 而且在输出的字符前面占了四个空行（由于\n）
// 重复使用 := 定义同名变量的话，后一个会覆盖前一个的值。并不会引起错误

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

// {{}}包围起来的部分就是go语法的内容，因此内部变量定义也必须使用 :=
// 变量前必须有 $
// 变量声明之后再修改只需使用等号 =