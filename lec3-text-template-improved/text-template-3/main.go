package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// 模板中的等式与不等式
// 模板中也能进行等式与不等式的判断，为更复杂的条件判断提供必要支持
// 用于条件判断的函数主要有以下6个：
// eq arg1 arg2 : 当 arg1 == arg2 成立时返回true, 否则 false		// equal
// ne arg1 arg2 : 当 arg1 != arg2 成立时返回true, 否则 false		// not equal
// lt arg1 arg2 : 当 arg1 < arg2 成立时返回true, 否则 false		// less than
// le arg1 arg2 : 当 arg1 <= arg2 成立时返回true, 否则 false		// less than or equal
// gt arg1 arg2 : 当 arg1 > arg2 成立时返回true, 否则 false		// greater than
// ge arg1 arg2 : 当 arg1 >= arg2 成立时返回true, 否则 false		// greater than or equal
//

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 创建模板对象并解析模板内容
		tmpl, err := template.New("test").Parse(`
{{$name1 := "alice"}}
{{$name2 := "bob"}}
{{$age1 := 21}}
{{$age2 := 25}}

{{if eq $age1 $age2}}
	两人年龄相同
{{else}}
	两人年龄不相同
{{end}}

{{if ne $name1 $name2}}
	两人名字不同
{{end}}

{{if gt $age1 $age2}}
	alice 年龄较大
{{else}}
	bob 年龄较大
{{end}}
`)

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

// 函数的调用一般是带括号的，只不过 text/template 可以在语法上省略这个括号