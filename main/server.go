package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	Models "go-micro-gin-gateway/models"
	"go-micro-gin-gateway/service"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	server := micro.NewService(
		micro.Name("user.server"),
		micro.Address("0.0.0.0:9000"),
		micro.Registry(reg),
	)
	Models.RegisterUserCommonServiceHandler(server.Server(), new(service.UserService))
	server.Init()
	server.Run()
}
