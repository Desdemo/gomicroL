package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	"time"
)

func main()  {
	// 连接consul
	cr := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// 循环获取
	for  {
		service, err := cr.GetService("userService")
		if err != nil{
			log.Fatalln(err)
		}

		next := selector.Random(service)
		svc, err := next()
		if err != nil{
			log.Fatalln(err)
		}
		fmt.Println("[测试输出]", svc.Address)
		time.Sleep(time.Second * 1)
	}
}
