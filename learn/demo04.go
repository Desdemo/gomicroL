package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
)

func main()  {
	// 连接到consul
	cr := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// 根据service nam 获取服务
	services, err := cr.GetService("productService")
	if err != nil {
		log.Fatalln(err)
	}
	// 使用random随机获取一个实例
	next := selector.Random(services)
	svc, err := next()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("[测试输出]",svc.Id, svc.Address, svc.Metadata)

}
