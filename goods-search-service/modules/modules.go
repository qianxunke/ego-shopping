package modules

import (
	"ego-goods-search-service/modules/product_es"
	productEsEervice "ego-goods-search-service/modules/product_es/service"

	productProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	"google.golang.org/grpc"
	"log"
)

func Init() {
	product_es.Init()
}

func RegisterService(s *grpc.Server) {
	//brand
	productEs, err := productEsEervice.GetService()
	if err != nil {
		log.Fatal("brand service init error")
		return
	}
	productProto.RegisterProductHandlerServer(s, productEs)

}
