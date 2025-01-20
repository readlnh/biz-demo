package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/readlnh/biz-demo/gomall/demo/demo_thrift/kitex_gen/api"
	"github.com/readlnh/biz-demo/gomall/demo/demo_thrift/kitex_gen/api/echo"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	c, err := echo.NewClient("demo_thrift", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	res, err := c.Echo(context.TODO(), &api.Request{Message: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", res)
}
