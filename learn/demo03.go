package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"net/http"
	_ "net/http"

	_ "github.com/gin-gonic/gin"
	_ "github.com/micro/go-micro/registry"
	_ "github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)


func main()  {
	// 添加consul地址
	cr := consul.NewRegistry(
			registry.Addrs("127.0.0.1:8500"))
	// 使用gin作为router
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "你好 打工人")
	})
	// 初始化go micro
	server := web.NewService(
			web.Address(":8081"),
			web.Name("productService"),
			web.Registry(cr),
			web.Metadata(map[string]string{"protocol":"http"}),
			web.Handler(router))

	_ = server.Run()
}
