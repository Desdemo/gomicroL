package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"net/http"
)

func main()  {
	// 1、添加consul地址
	cr := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// 2、使用gin作为router
	router := gin.Default()
	router.GET("/user", func(context *gin.Context) {
		context.String(http.StatusOK, "你好， 二号打工人")
	})
	// 3、初始化go micro
	server := web.NewService(
			web.Name("userService"),
			web.Registry(cr),
			web.Address(":8085"),
			web.Metadata(map[string]string{"protocol":"http"}),
			web.Handler(router))

	server.Init()

	_ = server.Run()
}
