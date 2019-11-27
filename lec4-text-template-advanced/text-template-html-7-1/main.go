package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
)

// 反转义
// 渲染 html时正确的姿势是使用html/template，它会对可疑内容进行转义，这是一个优点，
// 但是某种角度讲这也是缺点。
// 有的时候确实需要动态的生成html内容然后作为变量通过模板引擎进行渲染。
// 这时可以通过模板函数，将我们确信安全的文本转换为一个特殊类型 template.HTML
// 这样模板引擎就不会对其进行转义

// 有的时候确实需要将用户输入的内容渲染为HTML格式，怎么才能将任意文本安全的渲染为html且避免跨站脚本攻击呢？
// 有人开源了 bluemonkey 工具包， 这个包可以帮助我们渲染HTML时过滤掉所有潜在的不安全内容，而非无脑对所有字符转义
// 示例如下

func main() {

	// 创建 bluemonday的一个Policy对象
	p := bluemonday.UGCPolicy()

	// 创建模板
	tmpl := template.New("test")

	// 添加模板函数
	tmpl = tmpl.Funcs(template.FuncMap{
		"sanitize": func(s string) template.HTML {
			return template.HTML(p.Sanitize(s))		// 使用 p.Sanitize(s)先过滤
		},
	})

	// 同样的还有像 template.CSS/JS等其他类型，作用类似

	// 解析模板内容
	tmpl, err := tmpl.Parse(`
<html>
<body>
	<h2>Heading 2</h2>
	<p>{{.content | sanitize}}</p>
<body>
</html>
`)	// 这里是使用管道，将 .content内容传给 sanitize函数

	if err != nil {
		fmt.Printf("Parse: %s", err)
		return
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 调用模板对象的渲染方法
		err = tmpl.Execute(writer, map[string]interface{}{
			"content": `<a onblur="alert(secret)" href="http://www.baidu.com">Baidu</a>`,
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