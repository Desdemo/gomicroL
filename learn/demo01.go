package main

import (
	"github.com/micro/go-micro/web"
	"net/http"
)

func main()  {
	server := web.NewService(web.Address(":8081"))
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello go micro"))
	})

	_ = server.Run()
}
