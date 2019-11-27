package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// 反转义
// 渲染 html时正确的姿势是使用html/template，它会对可疑内容进行转义，这是一个优点，
// 但是某种角度讲这也是缺点。
// 有的时候确实需要动态的生成html内容然后作为变量通过模板引擎进行渲染。
// 这时可以通过模板函数，将我们确信安全的文本转换为一个特殊类型 template.HTML
// 这样模板引擎就不会对其进行转义
// 下面是示例

func main() {

	// 创建模板
	tmpl := template.New("test")

	// 添加模板函数
	tmpl = tmpl.Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})

	// 同样的还有像 template.CSS/JS等其他类型，作用类似

	// 解析模板内容
	tmpl, err := tmpl.Parse(`
<html>
<body>
	<h2>Heading 2</h2>
	<p>{{.content | safe}}</p>
<body>
</html>
`)	// 这里是使用管道，将 .content内容传给 safe函数

	if err != nil {
		fmt.Printf("Parse: %s", err)
		return
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 调用模板对象的渲染方法
		err = tmpl.Execute(writer, map[string]interface{}{
			"content": "<b>Hello World</b>",
		})
		if err != nil {
			fmt.Fprintf(writer, "Execute: %s", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
