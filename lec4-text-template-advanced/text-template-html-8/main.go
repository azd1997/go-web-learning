package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// 修改分隔符
// 标准库模板引擎使用的双层花括号{{}}和许多流行的前端框架有冲突（如VueJS、AngularJS）
// 所以必要时需要修改这个文本分隔符

func main() {

	// 创建模板
	tmpl := template.New("test")

	// 修改分隔符为 "[[""]]"
	tmpl = tmpl.Delims("[[", "]]")

	// 解析模板内容
	tmpl, err := tmpl.Parse(`
[[.content]]
`)	// 这里是使用管道，将 .content内容传给 sanitize函数

	if err != nil {
		fmt.Printf("Parse: %s", err)
		return
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 调用模板对象的渲染方法
		err = tmpl.Execute(writer, map[string]interface{}{
			"content": "Hello World",
		})		// 这里渲染了一个超链接到模板中
		if err != nil {
			fmt.Fprintf(writer, "Execute: %s", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// 从浏览器我们得到一个正确的文本超链接
// 再从curl url来看下究竟发生了什么
// (base) eiger@eiger-ThinkPad-X1-Carbon-3rd:~/gopath-default/src/github.com/azd1997/go-web-learning$ curl http://localhost:8000
//
//<html>
//<body>
//        <h2>Heading 2</h2>
//        <p><a href="http://www.baidu.com" rel="nofollow">Baidu</a></p>
//<body>
//</html>

// 可见 onblur="alert(secret)" 被过滤掉了