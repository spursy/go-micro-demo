package main

import (
	"context"
	"fmt"

	proto "go_micro_demo/proto"
	micro "github.com/micro/go-micro"
)

// Hello struct
type Hello struct{}

// Ping func
func (h *Hello) Ping(ctx context.Context, req *proto.Request, res *proto.Response) error {
	res.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("hellooo"),
	)
	service.Init()
	proto.RegisterHelloHandler(service.Server(), new(Hello))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}