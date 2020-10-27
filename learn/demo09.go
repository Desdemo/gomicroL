package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
	"log"
)

func main() {
	cr := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	myselector := selector.NewSelector(
		selector.Registry(cr),
		selector.SetStrategy(selector.RoundRobin))

	resp, err := callByGoPlugin(myselector)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("[服务调用结果]", resp)
}

func callByGoPlugin(s selector.Selector) (map[string]interface{}, error) {
	gopluginClient := http.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)

	req := gopluginClient.NewRequest("dataService", "/v1/list", map[string]interface{}{"size": 6})

	var resp map[string]interface{}
	err := gopluginClient.Call(context.Background(), req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
