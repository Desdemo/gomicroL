package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"net/http"


)

func main()  {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK,"早上好 打工人")
	})

	server := web.NewService(
		web.Address(":8081"),
		web.Metadata(map[string]string{"protol":"http"}),
		web.Handler(r))
	_ = server.Run()
}
