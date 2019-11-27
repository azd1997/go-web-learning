package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// text/template 与 html/template 的关联与区别
// 如本例所示，使用text/template也可以实现html的渲染
// 按照官方的说法，text/template是对text/template的一层封装，并在此基础上专注安全保障。
// 作为使用者，最直观的变化就是所有的文本变量进行了转义处理

func main() {

	// 创建模板
	tmpl := template.New("test")

	// 解析模板内容
	tmpl, err := tmpl.Parse(`
<html>
<body>
	<h2>Heading 2</h2>
	<p>this is a paragraph</p>
<body>
</html>
`)
	if err != nil {
		fmt.Printf("Parse: %s", err)
		return
	}

	// 打印下这个文件内容进行调试
	//file, err1 := os.Open("./lec4-text-template-advanced/text-template-5/test.tmpl")
	//data, err2 := ioutil.ReadAll(file)
	//if err1 != nil || err2 != nil {
	//	fmt.Println("err1: ", err1, "\n", "err2: ", err2)
	//} else {
	//	fmt.Print(string(data))
	//}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
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
