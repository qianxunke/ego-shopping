package main

import (
	"ego-goods-search-service/modules"
	_ "github.com/qianxunke/ego-shopping/ego-plugins/elasticsearch"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	//初始化
	modules.Init()
	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("err: v", err)
		return
	}

	s := grpc.NewServer()

	modules.RegisterService(s)

	reflection.Register(s)

	log.Printf("Server started at port %s", 9000)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
