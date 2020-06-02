package main

import (
	"ego-inventory-api/handler"
	"github.com/gin-gonic/gin"
	"github.com/qianxunke/ego-shopping/ego-common-utils/api_common/wrapper"
	"google.golang.org/grpc"
	"log"
)

func main() {
	con, err := grpc.Dial("localhost:8091", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer con.Close()
	router := gin.Default()
	router.Use(wrapper.AuthWrapper)
	handler.RegiserRouter(con, router)
	_ = router.Run(":8090")
}
