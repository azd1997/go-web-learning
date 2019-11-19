package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// 使用map作为模板根对象
// 和之前定义一个类型来作为根对象，这种方式更为灵活

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
`)

		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}

		// 获取 URL 参数的值
		sku := request.URL.Query().Get("sku")
		name := request.URL.Query().Get("name")

		unitPrice, _ := strconv.ParseFloat(request.URL.Query().Get("unitPrice"), 64)
		quantity, _ := strconv.ParseInt(request.URL.Query().Get("quantity"), 10, 64)

		// 调用模板对象的渲染方法。 	创建一个map[string]interface{}作为根对象
		err = tmpl.Execute(writer, map[string]interface{}{
			"SKU": sku,
			"Name": name,
			"UnitPrice": unitPrice,
			"Quantity": quantity,
		})
		if err != nil {
			fmt.Fprintf(writer, "Parse: %v", err)
			return
		}
	})

	log.Println("Starting HTTP Server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}