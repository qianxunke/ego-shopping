package main

import (
	"ego-user-service/modules/user_info"
	"ego-user-service/modules/user_info/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("err: v", err)
		return
	}

	s := grpc.NewServer()

	user_info.RegisterUserInfoServer(s, &service.UserInfoService{})

	reflection.Register(s)

	log.Printf("Server started at port %s", 8080)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
