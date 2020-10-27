package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"net/http"
	"strconv"
)

func main() {
	cr := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.POST("list", func(context *gin.Context) {
			var req ProdRequest
			if err := context.Bind(&req); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
					"data": "模型绑定失败",
				})
				context.Abort()
				return
			}

			context.JSON(http.StatusOK, gin.H{
				"data": NewProductList(req.Size),
			})
		})
	}

	service := web.NewService(
		web.Name("dataService"),
		web.Registry(cr),
		web.Address(":8087"),
		web.Handler(router),
		web.Metadata(map[string]string{"protocol": "http"}))

	_ = service.Init()

	_ = service.Run()
}

type ProdRequest struct {
	Size int `json:"size"`
}

type Product struct {
	Id   int
	Name string
}

func NewProduct(id int, name string) *Product {
	return &Product{
		Id:   id,
		Name: name,
	}
}

func NewProductList(count int) []*Product {
	products := make([]*Product, 0)
	for i := 0; i < count; i++ {
		products = append(products, NewProduct(i+1, "productName"+strconv.Itoa(i+1)))
	}
	return products
}
