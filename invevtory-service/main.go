package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/qianxunke/ego-shopping/ego-plugins/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"inventory-service/modules"
	"log"
	"net"
)

func main() {
	modules.Init()
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("err: v", err)
		return
	}

	s := grpc.NewServer()

	modules.RegisterService(s)

	reflection.Register(s)

	log.Printf("Server started at port %s", 8081)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
