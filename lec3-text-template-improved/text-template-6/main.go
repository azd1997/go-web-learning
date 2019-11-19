package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// 模板中空白符号处理
// 使用 {{- 可以去除模板左侧的所有空白符号， -}}去除所有右侧空白符号
// 注意 - 短横线要与模板内其他内容隔开来，不然模板引擎会误认为是表达式一部分

type Inventory struct {
	SKU string
	Name string
	UnitPrice float64
	Quantity int64
}

func (i *Inventory) Subtotal() float64 {
	return i.UnitPrice * float64(i.Quantity)
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 创建模板对象并解析模板内容		// 注意注释行仍会存在（因为有\n），但注释不显示
		tmpl, err := template.New("test").Parse(`{{/* 打印参数值 */}}
不去除空白符号的话，下面会有一个空行：
{{with .Inventory}}
	SKU: {{.SKU}}
	Name: {{.Name}}
	UnitPrice: {{.UnitPrice}}
	Quantity: {{.Quantity}}
{{end}}
这句话上面也有个空行


去除空白符号的话，这句话会和变量输出紧邻：
{{- with .Inventory}}
	SKU: {{.SKU}}
	Name: {{.Name}}
	UnitPrice: {{.UnitPrice}}
	Quantity: {{.Quantity}}
{{- end}}
现在这句话上面没空行了（其实都是去掉了左侧的\n）
`)

		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}

		// 调用模板对象的渲染方法。 	创建一个map[string]interface{}作为根对象
		err = tmpl.Execute(writer, map[string]interface{}{
			"Inventory": Inventory{
				SKU:       "110000",
				Name:      "phone",
				UnitPrice: 7888.2,
				Quantity:  20,
			},
		})
		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}
	})

	log.Println("Starting HTTP Server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}