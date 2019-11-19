package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// 在模板中渲染复杂变量



type Inventory struct {
	SKU string
	Name string
	UnitPrice float64
	Quantity int64
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 创建模板对象并解析模板内容
		tmpl, err := template.New("test").Parse(`Inventory
SKU: {{.SKU}}
Name: {{.Name}}
UnitPrice: {{.UnitPrice}}
Quantity: {{.Quantity}}
`)

		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}

		// 获取 URL 参数的值, 创建Inventory实例
		inventory := &Inventory{
			SKU: request.URL.Query().Get("sku"),
			Name: request.URL.Query().Get("name"),
		}
		inventory.UnitPrice, _ = strconv.ParseFloat(request.URL.Query().Get("unitPrice"), 64)
		inventory.Quantity, _ = strconv.ParseInt(request.URL.Query().Get("quantity"), 10, 64)

		// 调用模板对象的渲染方法
		err = tmpl.Execute(writer, inventory)		// 将变量val渲染到io.writer
		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}
	})

	log.Println("Starting HTTP Server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}