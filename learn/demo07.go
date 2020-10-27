package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// 客户端API 请求
func main()  {
	// 连接consul
	cr := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// 获取微服务列表
	service, err := cr.GetService("userService")
	if err != nil{
		log.Fatalln(err)
	}
	// 使用random随机读取一个实例
	next := selector.Random(service)
	svc, err := next()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("[测试输出]", svc.Address,"/user", nil)
	// 请求获取到服务的api 方法
	resp, err := RequestApi(http.MethodGet, svc.Address, "/user", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("[请求API结果]", resp)
}

func RequestApi(method string, host string, path string, body io.Reader) (string, error)  {
	if !strings.HasPrefix(host,"http://") && !strings.HasPrefix(host, "https://") {
		host = "http://" + host
	}
	req, _ := http.NewRequest(method, host+path, body)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer  res.Body.Close()

	buff, err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Fatalln(err)
	}

	return string(buff), nil
}
