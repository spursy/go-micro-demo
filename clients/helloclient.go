package main

import (
	"context"
	"fmt"

	proto "go_micro_demo/proto"
	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("hello.client"))
	service.Init()
	helloService := proto.NewHelloService("hellooo", service.Client())
	res,err := helloService.Ping(context.TODO(), &proto.Request{Name: "World ^_^"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Msg)
}