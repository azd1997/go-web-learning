package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// 在模板中调用结构的方法


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
		// 创建模板对象并解析模板内容
		tmpl, err := template.New("test").Parse(`Inventory
SKU: {{.SKU}}
Name: {{.Name}}
UnitPrice: {{.UnitPrice}}
Quantity: {{.Quantity}}
Subtotal: {{.Subtotal}}
`)		// 注意： 直接调用方法名，template模板引擎会自动识别调用对象的具体类型

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