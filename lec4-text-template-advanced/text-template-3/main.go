package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// 模板中的管道操作
// 前一个函数的输出可以通过管道传给下一个函数作为输入
// 哪怕是多输入多输出函数也可以使用管道

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 创建模板
		tmpl := template.New("test")

		// 添加自定义函数
		tmpl = tmpl.Funcs(template.FuncMap{
			// 单输入单输出函数
			"add1": func(a int) int {
				return a+2
			},

			// 单输入单输出好理解，但令人困惑的是多输入多输出情况，管道传来的参数是作为当前函数入参的哪一个呢？是靠左，中间，还是靠右？

			// 多输入单输出函数
			"add2": func(a, b int) int {
				return a+b
			},

			// 多输入多输出函数
			"add3": func(a int) int {
				return a+2
			},

			// minus1 和 minus2 用来比较管道传来的是默认给最左还是最右的参数
			"minus1": func(a, b int) int {
				return a-b
			},
			"minus2": func(a, b int) int {
				return b-a
			},
			// 经过测试，管道传过去的参数，默认是传给靠右的参数
			//result4: 1
			//result5: -5


		})

		// 解析模板内容
		tmpl, err := tmpl.Parse(`
result1: {{add1 0 | add1 | add1}}
result2: {{add2 1 3 | add2 2 | add2 2}}
result3:
result4: {{minus1 4 2 | minus1 3}}
result5: {{minus2 4 2 | minus2 3}}
`)
		// result1: 管道操作 add2 0 = 2; 传给下一个add2 得到4； 传给下一个，得到6
		// result2: 8



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

