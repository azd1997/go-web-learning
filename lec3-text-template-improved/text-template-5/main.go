package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// 在模板中使用语境操作（with语句）

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
不使用with语句的话，需要写成：
SKU: {{.Inventory.SKU}}
Name: {{.Inventory.Name}}
UnitPrice: {{.Inventory.UnitPrice}}
Quantity: {{.Inventory.Quantity}}

当需要使用连续 . 操作后，语句变得冗长不易读

使用with语境，写成这样：
{{with .Inventory}}
	SKU: {{.SKU}}
	Name: {{.Name}}
	UnitPrice: {{.UnitPrice}}
	Quantity: {{.Quantity}}
{{end}}

模板更加简洁易懂了。
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