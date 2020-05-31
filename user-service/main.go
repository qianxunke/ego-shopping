package main

import (
	"ego-user-service/clients"
	"ego-user-service/modules"
	"ego-user-service/modules/user_info/service"
	"github.com/qianxunke/ego-shopping/ego-common-protos/out/user_info"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	//初始化
	clients.Init()
	modules.Init()
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("err: v", err)
		return
	}

	s := grpc.NewServer()
	userInfoService, err := service.GetService()
	if err != nil {
		log.Fatal("userInfoService init error")
		return
	}
	user_info.RegisterUserInfoServer(s, userInfoService)

	reflection.Register(s)

	log.Printf("Server started at port %s", 8080)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
