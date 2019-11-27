package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// 从本地文件加载模板

func main() {

	// 这里将创建模板和解析提到前面是因为，这样在用户请求过来的时候就不用再走一遍这些，节省时间

	// 创建模板
	tmpl := template.New("test")

	// 解析模板内容
	tmpl, err := tmpl.ParseFiles("./lec4-text-template-advanced/text-template-5/test.tmpl")	// 这个文件后缀名无所谓，一般标准库的模板后缀名起.tmpl或.tpl
	// 这里因为 go run main.go的当前目录是项目根目录，所以模板路径填的是这个而不是 ./test.tmpl
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
		err = tmpl.Execute(writer, map[string]interface{}{
			"names": []string{"Alice", "Bob", "Cindy", "David"},
		})
		if err != nil {
			fmt.Fprintf(writer, "Execute: %s", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

//(base) eiger@eiger-ThinkPad-X1-Carbon-3rd:~/gopath-default/src/github.com/azd1997/go-web-learning$ curl http://localhost:8000
//Execute: template: test: "test" is an incomplete or empty template

// ParseFiles()可以传任意多个模板文件，那么在handle函数中如何指定 Execute(渲染)哪一个模板呢？
// 用 ExecuteTemplate(w io,writer, templateFileName string, data interface{}) 参数2指定加载哪一个模板文件