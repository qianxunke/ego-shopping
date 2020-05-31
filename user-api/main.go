package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"user-api/handler"
)

func main() {
	con, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer con.Close()
	router := gin.Default()
	router.Use(handler.AuthWrapper)
	handler.RegiserRouter(con, router)
	_ = router.Run(":8081")
}
