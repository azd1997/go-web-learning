package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// 模板中使用条件判断（if语句）
// 模板引擎中可以像程序代码那样进行基本的逻辑控制。

// 这里设计一个除法服务，当客户端给定的参数不合理时要提示客户端参数错误，这就用上了if

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 创建模板对象并解析模板内容
		tmpl, err := template.New("test").Parse(`
{{if .yIsZero}}
	除数不能为0
{{else}}
	{{.result}}
{{end}}
`)		// 注意 if后的条件语句必须返回一个bool值
// if ... else ... end

		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}

		// 获取 URL 查询参数的值
		x, _ := strconv.ParseInt(request.URL.Query().Get("x"), 10, 64)
		y, _ := strconv.ParseInt(request.URL.Query().Get("y"), 10, 64)

		// 当y不为0时进行除法运算
		yIsZero := y==0
		result := 0.0	// float64
		if !yIsZero {
			result = float64(x)/float64(y)
		}


		// 调用模板对象的渲染方法
		err = tmpl.Execute(writer, map[string]interface{}{
			"yIsZero": yIsZero,
			"result": result,
		})
		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}
	})

	log.Println("Starting HTTP Server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

// 实际的开发中，if经常使用，然后根据给定的条件判断渲染出不同的内容