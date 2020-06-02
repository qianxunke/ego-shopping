package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxunke/ego-shopping/ego-common-utils/api_common/wrapper"
	"google.golang.org/grpc"
	"log"
	"user-api/handler"
)

func main() {
	con, err := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer con.Close()
	router := gin.Default()
	router.Use(wrapper.AuthWrapper)
	handler.RegiserRouter(con, router)
	_ = router.Run(":8080")
}
